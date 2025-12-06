package utils

import (
	"math/rand"
	"time"
)

func GenerateVerificationCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}
