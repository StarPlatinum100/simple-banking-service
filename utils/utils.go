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
