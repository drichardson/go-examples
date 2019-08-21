package main

import (
	"github.com/pascaldekloe/jwt"
	"log"
	"time"
)

func main() {
	// tokens generated at https://jwt.io/

	// no expiration time... fails claims.Expires == nil check below
	// token := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.-eI_wZ-C8gYVg0VEPccXnjGzbq1D7bP56s3vw7ZNGxM")

	token := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTY2MzUyNDM2LCJleHAiOjE1NjYzNTI0MzV9.lH6ZoSuWfBaWY_S7BjaNAk-vwQvQQRSQ4j3rYQRUp80")
	secret := []byte("my secret value")

	claims, err := jwt.HMACCheck(token, secret)
	if err != nil {
		log.Fatal("HMACCheck failed. ", err)
	}

	log.Printf("Claims: %v", claims)
	log.Printf("Claim expires at %v", claims.Expires)

	if claims.Expires == nil {
		log.Fatal("Must have expiration time (my own requirement, not a jwt requirement)")
	}

	// not yet expired
	fakeNow := time.Unix(1566352400, 0)
	// too late, expired
	// fakeNow := time.Unix(1566352500, 0)
	if !claims.Valid(fakeNow) {
		log.Fatal("credential time constraints exceeded")
	}

	log.Println("OK")
}
