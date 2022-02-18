package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomStringAlphabet(n int) string {
	var stringBuilder strings.Builder

	lenAlpha := len(alphabet)
	for i := 0; i < n; i++ {
		char := alphabet[rand.Intn(lenAlpha)]
		stringBuilder.WriteByte(char)
	}

	return stringBuilder.String()
}

func RandomUserName() string {
	return RandomStringAlphabet(10)
}

func RandomRole() string {
	roles := []string{"admin", "scraper"}
	return roles[rand.Intn(len(roles))]
}
