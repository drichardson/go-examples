package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

var generate = flag.Bool("new", false, "Generate a new token.")
var validate = flag.String("validate", "", "Validate a token.")
var key = flag.String("key", "12345", "The key to use.")

func main() {
	flag.Parse()

	if *generate {
		expire := time.Now().Add(10 * time.Second).Unix()
		b := bytes.NewBuffer(nil)
		e := json.NewEncoder(b)
		if err := e.Encode(expire); err != nil {
			fmt.Println("Error encoding expiration.", err)
			os.Exit(1)
		}
		h := hmac.New(sha256.New, []byte(*key))
		if _, err := h.Write(b.Bytes()); err != nil {
			fmt.Println("Error wriring bytes to hmac.", err)
			os.Exit(1)
		}
		if _, err := b.Write(h.Sum(nil)); err != nil {
			fmt.Println("Error appending hmac to bytes.", err)
			os.Exit(1)
		}
		fmt.Println(hex.EncodeToString(b.Bytes()))
		os.Exit(0)
	}

	if *validate != "" {
		b, err := hex.DecodeString(*validate)
		if err != nil {
			fmt.Println("Failed to decode hex string", err)
			os.Exit(1)
		}
		if len(b) < sha256.Size {
			fmt.Println("Invalid size")
			os.Exit(1)
		}
		givenHMAC := b[len(b)-sha256.Size:]
		givenJsonBytes := b[0 : len(b)-sha256.Size]
		computedH := hmac.New(sha256.New, []byte(*key))
		if _, err := computedH.Write(givenJsonBytes); err != nil {
			fmt.Println("Error writing json bytes to hmac.", err)
			os.Exit(1)
		}
		computedHMAC := computedH.Sum(nil)
		if !hmac.Equal(givenHMAC, computedHMAC) {
			fmt.Println("HMAC doesn't match")
			os.Exit(1)
		}

		byteReader := bytes.NewReader(givenJsonBytes)
		decoder := json.NewDecoder(byteReader)
		var expiry int64
		if err := decoder.Decode(&expiry); err != nil {
			fmt.Println("Failed to decode time.", err)
			os.Exit(1)
		}
		if time.Now().Unix() > expiry {
			fmt.Println("Time expired")
			os.Exit(1)
		}

		fmt.Println("Valid Token")
		os.Exit(0)
	}

	fmt.Println("No command given")
	flag.PrintDefaults()
	os.Exit(1)
}
