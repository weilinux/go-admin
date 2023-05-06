package controller

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func generatedHash(password []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("generated hash: %v", string(hash))
	return hash
}

func compareHash(password []byte, inputPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(password, inputPassword)
	if err != nil {
		// password does not match
		return false
	}
	return true
}
