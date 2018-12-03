package helper

import "math/rand"

const (
	textForRandomUsername = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	textForRandomPassword = "0123456789@#!_-+&$M?"
)

func RandomUsername() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = textForRandomUsername[rand.Intn(len(textForRandomUsername))]
	}
	return string(b)
}

func RandomPassword() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = textForRandomPassword[rand.Intn(len(textForRandomPassword))]
	}
	return string(b)
}
