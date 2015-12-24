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

	h := hmac.New(sha256.New, []byte("key"))
	h.Write([]byte("The quick brown fox jumps over the lazy dog"))
	fmt.Println("wikipedia check: f7bc83f430538424b13298e6aa6fb143ef4d59a14946175997479dbc2d1a3cd8=?", hex.EncodeToString(h.Sum(nil)))

	h2 := hmac.New(sha256.New, []byte(""))
	h2.Write([]byte(""))
	fmt.Println("wikipedia check: b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad=?", hex.EncodeToString(h2.Sum(nil)))

}

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
