package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `hello`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("main", pass)
	fmt.Println("bcrypt", bs)

	loginPword := `hello`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword))
	if err != nil {
		fmt.Println("COMPARE ERROR", err)
	}
	fmt.Println("pass")

}
