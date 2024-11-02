package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHash(value string) string {
	hash := sha256.New()
	hash.Write([]byte(value))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

