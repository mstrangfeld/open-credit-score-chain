package records_test

import (
	"bytes"
	"testing"

	ecies "github.com/ecies/go"
	"github.com/mstrangfeld/open-credit-score-chain/backend/records"
)

func TestEncryption(t *testing.T) {
	csr := &records.CreditScoreRecord{
		Reason: "Something",
		Score:  100,
	}
	priv, err := ecies.GenerateKey()
	if err != nil {
		t.Error(err)
	}
	pub := priv.PublicKey.Hex(true)
	encrypted, err := csr.Encrypt(pub)
	if err != nil {
		t.Error(err)
	}
	decrypted, err := encrypted.Decrypt(priv.Hex())
	if err != nil {
		t.Error(err)
	}
	if decrypted.Reason != csr.Reason {
		t.Error("Decrypted reason does not match")
	}
	if decrypted.Score != csr.Score {
		t.Error("Decrypted score does not match")
	}
	regeneratedHash := records.GeneratePublicScoreHash(decrypted.Score, decrypted.Nonce)
	if !bytes.Equal(regeneratedHash, decrypted.ScoreHash) {
		t.Error("Decrypted hash does not match")
	}
}
