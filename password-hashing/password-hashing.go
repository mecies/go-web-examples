package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {

	originalPassword := "secret"
	password := originalPassword
	oneOrTwo := rand.Intn(3-1) + 1

	if oneOrTwo == 1 {
		password = "different"
	}

	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:", hash)

	match := CheckPasswordHash(originalPassword, hash)

	fmt.Println("Match:   ", match)
}
