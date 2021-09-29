package mimc

import (
	gohash "hash"

	"github.com/consensys/gnark-crypto/hash"
)

const (
	mimcSeed = "seed"
	mimcHash = hash.MIMC_BN254
)

// NewMimcHash returns a new hash function for hashing the score with the nonce
func NewMimcHash() gohash.Hash {
	return mimcHash.New(mimcSeed)
}
