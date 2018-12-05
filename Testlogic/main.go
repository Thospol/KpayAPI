package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// a := StringWithMerchantset("jdfgnjtedkgfkntm")
	// fmt.Print(a)
	fmt.Println(time.Now().Format("02-01-2006"))
}

func StringWithMerchantset(Merchantset string) string {

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 10)
	for i := range b {
		b[i] = Merchantset[seededRand.Intn(len(Merchantset))]
	}
	return string(b)
}

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
