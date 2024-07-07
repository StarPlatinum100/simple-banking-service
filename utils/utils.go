package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateAccountNumber() string {

	rand.Seed(time.Now().UnixNano())
	accountNumber := ""

	for i := 0; i < 10; i++ {
		accountNumber += fmt.Sprintf("%d", rand.Intn(10))
	}

	return accountNumber
}

func GenerateRandomString() string {
	rand.Seed(time.Now().UnixNano())

	const charset = "0123456789"
	const length = 20

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
