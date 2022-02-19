package util

import (
	"errors"
	"fmt"
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

func RandomBool() bool {
	if rand.Intn(2) == 0 {
		return false
	}
	return true
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
	roles := GetAllRoles()
	return roles[rand.Intn(len(roles))]
}

func RandomBlockchainName() string {
	return RandomStringAlphabet(6)
}

func RandomEventName() string {
	return RandomStringAlphabet(6)
}

func RandomEventFieldName() string {
	return RandomStringAlphabet(4)
}

func RandomEvenFieldType() string {
	allEventFieldTypes := GetAllEventFieldTypes()
	return allEventFieldTypes[rand.Intn(len(allEventFieldTypes))]
}

func RandomEventFieldValue(eventFieldType string) (string, error) {
	switch eventFieldType {
	case EventFieldString:
		return RandomStringAlphabet(10), nil
	case EventFieldNumber:
		return fmt.Sprintf("%g", rand.Float64()), nil
	case EventFieldBoolean:
		return fmt.Sprintf("%t", RandomBool()), nil // zero or one
	}
	return "", errors.New("unknown event field type")
}
