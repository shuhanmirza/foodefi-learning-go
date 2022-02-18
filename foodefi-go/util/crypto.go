package util

import "crypto/sha256"

func GetHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
