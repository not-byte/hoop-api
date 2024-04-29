package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GenerateHash(length int) string {
	hash := make([]rune, length)

	for i := range hash {
		hash[i] = characters[rand.Intn(len(characters))]
	}

	return string(hash)
}
