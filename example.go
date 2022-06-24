package main

import (
	"log"

	auth "github.com/shadowlion/firebaseauth-go/auth"
)

func main() {
	email := ""
	password := ""
	apiKey := ""

	AuthClient := auth.New(apiKey)
	resp, err := AuthClient.SignInWithPassword(email, password, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp.Email)
	// assert email == resp.Email
}
