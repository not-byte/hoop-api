package utils

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var numbers = []rune("0123456789")

func GenerateHash(length int) string {
	hash := make([]rune, length)

	for i := range hash {
		hash[i] = characters[rand.Intn(len(characters))]
	}

	return string(hash)
}

func GenerateToken() int8 {
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

func HandleCookieError(w http.ResponseWriter, err error, tokenType string) {
	if err == http.ErrNoCookie {
		http.Error(w, "Missing "+tokenType, http.StatusUnauthorized)
	} else {
		http.Error(w, "Invalid request for "+tokenType, http.StatusBadRequest)
	}
}
