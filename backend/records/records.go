package records

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	ecies "github.com/ecies/go"
	"github.com/mstrangfeld/open-credit-score-chain/backend/mimc"
	"github.com/rs/zerolog/log"
)

// Decrypted CreditScoreRecord
type CreditScoreRecord struct {
	Reason string
	Score  int
	// leave empty for encryption
	Nonce string
	// leave empty for encryption
	ScoreHash []byte
}

// EncrytedCreditScoreRecord is the encrypted credit score record.
type EncryptedCreditScoreRecord struct {
	Reason    string
	Score     string
	Nonce     string
	ScoreHash string
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GeneratePublicScoreHash(score int, nonce string) []byte {
	hash := mimc.NewMimcHash()
	var scoreElement, nonceElement fr.Element
	scoreElement.SetInterface(score)
	nonceElement.SetBytes([]byte(nonce))
	hash.Write(scoreElement.Marshal())
	hash.Write(nonceElement.Marshal())
	return hash.Sum(nil)
}

// Encrypt encrypts a credit score record using the given public key.
func (c *CreditScoreRecord) Encrypt(publicKeyHex string) (*EncryptedCreditScoreRecord, error) {
	pub, err := ecies.NewPublicKeyFromHex(publicKeyHex)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse public key")
		return nil, err
	}

	b, err := generateRandomBytes(8)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate random bytes")
		return nil, err
	}

	plainNonce := base64.URLEncoding.EncodeToString(b)
	scoreHash := GeneratePublicScoreHash(c.Score, plainNonce)

	encryptedNonce, err := ecies.Encrypt(pub, []byte(plainNonce))
	if err != nil {
		log.Error().Err(err).Msg("Failed to encrypt nonce")
		return nil, err
	}
	encryptedReason, err := ecies.Encrypt(pub, []byte(c.Reason))
	if err != nil {
		log.Error().Err(err).Msg("Failed to encrypt reason")
		return nil, err
	}
	encryptedScore, err := ecies.Encrypt(pub, big.NewInt(int64(c.Score)).Bytes())
	if err != nil {
		log.Error().Err(err).Msg("Failed to encrypt score")
		return nil, err
	}

	return &EncryptedCreditScoreRecord{
		Reason:    base64.URLEncoding.EncodeToString(encryptedReason),
		Score:     base64.URLEncoding.EncodeToString(encryptedScore),
		Nonce:     base64.URLEncoding.EncodeToString(encryptedNonce),
		ScoreHash: base64.URLEncoding.EncodeToString(scoreHash),
	}, nil

}

// Decrypt decrypts a credit score record using the given private key.
func (e *EncryptedCreditScoreRecord) Decrypt(privateKeyHex string) (*CreditScoreRecord, error) {
	priv, err := ecies.NewPrivateKeyFromHex(privateKeyHex)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse private key")
		return nil, err
	}
	decodedReason, err := base64.URLEncoding.DecodeString(e.Reason)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode reason")
		return nil, err
	}
	decryptedReason, err := ecies.Decrypt(priv, decodedReason)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decrypt reason")
		return nil, err
	}
	decodedScore, err := base64.URLEncoding.DecodeString(e.Score)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode score")
		return nil, err
	}
	decryptedScore, err := ecies.Decrypt(priv, decodedScore)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decrypt score")
		return nil, err
	}
	decodedNonce, err := base64.URLEncoding.DecodeString(e.Nonce)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode nonce")
		return nil, err
	}
	decryptedNonce, err := ecies.Decrypt(priv, decodedNonce)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decrypt nonce")
		return nil, err
	}
	decodedScoreHash, err := base64.URLEncoding.DecodeString(e.ScoreHash)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode score hash")
		return nil, err
	}
	return &CreditScoreRecord{
		Reason:    string(decryptedReason),
		Score:     int(new(big.Int).SetBytes(decryptedScore).Int64()),
		Nonce:     string(decryptedNonce),
		ScoreHash: decodedScoreHash,
	}, nil
}
