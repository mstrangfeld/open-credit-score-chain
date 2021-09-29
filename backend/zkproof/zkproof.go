package zkproof

import (
	"bytes"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
	"github.com/mstrangfeld/open-credit-score-chain/backend/records"
	"github.com/rs/zerolog/log"
)

const (
	batchSize = 10 // Number of records to process in one batch
)

// The zk Proof Circuit
type Circuit struct {
	// Private inputs
	Scores [batchSize]frontend.Variable
	Nonces [batchSize]frontend.Variable
	// Public inputs
	Hashes [batchSize]frontend.Variable `gnark:",public"`
	Score  frontend.Variable            `gnark:",public"`
}

// Calculate the sum of all the scores and proove that the sum is equal to the score
func (circuit *Circuit) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {
	mimc, _ := mimc.NewMiMC("seed", curveID, cs)
	output := cs.Constant(0)
	for i := 0; i < batchSize; i++ {
		mimc.Write(circuit.Scores[i])
		mimc.Write(circuit.Nonces[i])
		comp := cs.Select(cs.IsZero(circuit.Hashes[i]), circuit.Hashes[i], mimc.Sum())
		cs.AssertIsEqual(circuit.Hashes[i], comp)
		output = cs.Add(output, circuit.Scores[i])
		mimc.Reset()
	}
	cs.AssertIsEqual(circuit.Score, output)

	return nil
}

func calculateFinalScore(records []*records.CreditScoreRecord) int {
	finalScore := 0
	for i := 0; i < batchSize; i++ {
		if i < len(records) {
			finalScore += records[i].Score
		}
	}
	return finalScore
}

// CreateProofForRecords creates a proof for the given records
// Returns the proof, the verifying key and the final score
func CreateProofForRecords(records []*records.CreditScoreRecord) ([]byte, []byte, int, error) {
	var mimcCircuit Circuit

	r1cs, err := frontend.Compile(ecc.BN254, backend.GROTH16, &mimcCircuit)
	if err != nil {
		log.Error().Err(err).Msg("Failed to compile the circuit")
		return nil, nil, 0, err
	}

	// groth16 zkSNARK
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		log.Error().Err(err).Msg("Failed to setup the zkSNARK")
		return nil, nil, 0, err
	}

	finalScore := calculateFinalScore(records)

	var witness, publicWitness Circuit

	for i := 0; i < batchSize; i++ {
		if i < len(records) {
			witness.Scores[i].Assign(records[i].Score)
			witness.Nonces[i].Assign([]byte(records[i].Nonce))
			witness.Hashes[i].Assign(records[i].ScoreHash)
			publicWitness.Hashes[i].Assign(records[i].ScoreHash)
		} else {
			witness.Scores[i].Assign(0)
			witness.Nonces[i].Assign(0)
			witness.Hashes[i].Assign(0)
			publicWitness.Hashes[i].Assign(0)
		}
	}
	witness.Score.Assign(finalScore)
	publicWitness.Score.Assign(finalScore)

	proof, err := groth16.Prove(r1cs, pk, &witness)
	if err != nil {
		log.Error().Err(err).Msg("Failed to prove the circuit")
		return nil, nil, 0, err
	}

	err = groth16.Verify(proof, vk, &publicWitness)
	if err != nil {
		log.Error().Err(err).Msg("Failed to verify the circuit")
		return nil, nil, 0, err
	}

	log.Debug().Msg("Proof created successfully")

	rawProof := bytes.NewBuffer([]byte{})
	proof.WriteRawTo(rawProof)

	rawVk := bytes.NewBuffer([]byte{})
	vk.WriteRawTo(rawVk)

	return rawProof.Bytes(), rawVk.Bytes(), finalScore, nil
}

// VerifyProof verifies the given proof with the given verifying key and the final score
// Returns true if the proof is valid, false otherwise
func VerifyProof(rawProof []byte, rawVk []byte, hashes [][]byte, score int) (bool, error) {
	proof := groth16.NewProof(ecc.BN254)
	_, err := proof.ReadFrom(bytes.NewBuffer(rawProof))
	if err != nil {
		log.Error().Err(err).Msg("Failed to read the proof")
		return false, err
	}

	vk := groth16.NewVerifyingKey(ecc.BN254)
	_, err = vk.ReadFrom(bytes.NewBuffer(rawVk))
	if err != nil {
		log.Error().Err(err).Msg("Failed to read the verifying key")
		return false, err
	}

	var ownWitness Circuit
	for i := 0; i < batchSize; i++ {
		if i < len(hashes) {
			ownWitness.Hashes[i].Assign(hashes[i])
		} else {
			ownWitness.Hashes[i].Assign(0)
		}
	}
	ownWitness.Score.Assign(score)

	err = groth16.Verify(proof, vk, &ownWitness)
	if err != nil {
		log.Error().Err(err).Msg("Proof verification failed")
		return false, err
	}

	return true, nil
}
