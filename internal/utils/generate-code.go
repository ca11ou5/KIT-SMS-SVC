package utils

import (
	"math/rand"
	"time"
)

func GenerateCode() string {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, 6)
	for i := range code {
		code[i] = characters[rand.Intn(len(characters))]
	}

	return string(code)
}
