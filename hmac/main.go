package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("123456789abcdef")
	mac := hmac.New(sha256.New, key)
	message := "This is my test message."
	_, err := mac.Write([]byte(message))
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	sum := mac.Sum(nil)
	fmt.Println("message", message, "mac", hex.EncodeToString(sum))

	fmt.Println("checkmac", CheckMAC([]byte(message), sum, key))
	fmt.Println("checkmac (bad key)", CheckMAC([]byte(message), sum, key[:len(key)-1]))
	fmt.Println("checkmac (+1)", CheckMAC([]byte(message+"1"), sum, key))
	fmt.Println("checkmac (-1)", CheckMAC([]byte(message[:len(message)-1]), sum, key))
}

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
