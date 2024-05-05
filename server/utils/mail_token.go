package utils

import (
	"log"
	"math/rand"
	"strconv"
)

var numbers = []rune("0123456789")

func GenerateMailToken() int8 {
	hash := make([]rune, 6)

	for i := range hash {
		hash[i] = numbers[rand.Intn(len(numbers))]
	}

	token, err := strconv.Atoi(string(hash))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return int8(token)
}
