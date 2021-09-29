package zkproof_test

import (
	"encoding/base64"
	"testing"

	ecies "github.com/ecies/go"
	"github.com/mstrangfeld/open-credit-score-chain/backend/records"
	"github.com/mstrangfeld/open-credit-score-chain/backend/zkproof"
)

func TestProoving(t *testing.T) {
	priv, err := ecies.GenerateKey()
	if err != nil {
		t.Error(err)
	}
	pub := priv.PublicKey.Hex(true)
	recs := []*records.CreditScoreRecord{
		{
			Score:  1,
			Reason: "You're a bad guy",
		},
		{
			Score:  50,
			Reason: "You're an okay guy",
		},
		{
			Score:  100,
			Reason: "You're a good guy",
		},
	}
	encryptedRecs := []*records.EncryptedCreditScoreRecord{}
	decryptedRecs := []*records.CreditScoreRecord{}
	for _, record := range recs {
		enc, err := record.Encrypt(pub)
		if err != nil {
			t.Error(err)
		}
		encryptedRecs = append(encryptedRecs, enc)
		dec, err := enc.Decrypt(priv.Hex())
		if err != nil {
			t.Error(err)
		}
		decryptedRecs = append(decryptedRecs, dec)
	}

	proofData, vkData, score, err := zkproof.CreateProofForRecords(decryptedRecs)
	if err != nil {
		t.Error(err)
	}

	t.Log("Proof creation successful")

	hashes := [][]byte{}
	for _, record := range encryptedRecs {
		rawHash, err := base64.URLEncoding.DecodeString(record.ScoreHash)
		if err != nil {
			t.Error(err)
		}
		hashes = append(hashes, rawHash)
	}
	ok, err := zkproof.VerifyProof(proofData, vkData, hashes, score)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("Proof failed to verify")
	}
}
