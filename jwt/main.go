package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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

	log.Println("HMAC JWT OK")

	testRSA256()
}

func testRSA256() {
	// no expiration
	// token := []byte("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.POstGetfAytaZS82wHcjoTyoqhMyxXiWdR7Nn7A29DNSl0EiXLdwJ6xC6AfgZWF1bOsS_TuYI3OG85AmiExREkrS6tDfTQ2B3WXlrr-wp5AokiRbz3_oB4OxG-W9KcEEbDRcZc0nH3L7LzYptiy1PtAylQGxHTWZXtGz4ht0bAecBgmpdgXMguEIcoqPJ1n3pIWk_dUZegpqx0Lka21H6XxUTxiy8OcaarA8zdnPUnV6AmNP3ecFawIFYdvJB_cm-GvpCSbr8G8y_Mllj8f4x9nBH8pQux89_6gUY618iYv7tuPWBFfEbLxtF2pZS6YC1aSfLQxeNe8djT9YjpvRZA")

	// expires 1516239100
	token := []byte("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMiwiZXhwIjoxNTE2MjM5MTAwfQ.YVGcQGZqHenoh-ljWSMwxDWsxJqW-j6sw5ofW6E3NHATyiLmLSusUyy5zQm_nmYbBq8UWRv007Y9gQEcTgpw7_nHVLqegC3_Gz-t9coocQqDLrUMpU8iipp5cfJQTV35YLl8TRJMQqdaN5RvbDeWX6qmxNXuvgpcAI0DBkAw901YPe1Ko_pgQv3-USIh3snXZSXJ5tqkV32smopkf3i0UGc9g-fXeOVGgW_fM_UreULLGysq_faoNiT2Q7Om9FQISQ5cS88uRHvJywdwKtv7IGKEDmiFb9eQSoSfa_68Vk9fbh8d6dssCUcaigGA-E11VR8PX9FvpTA_YGn9yNWMww")

	publicKeyPEM := []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnzyis1ZjfNB0bBgKFMSv
vkTtwlvBsaJq7S5wA+kzeVOVpVWwkWdVha4s38XM/pa/yr47av7+z3VTmvDRyAHc
aT92whREFpLv9cj5lTeJSibyr/Mrm/YtjCZVWgaOYIhwrXwKLqPr/11inWsAkfIy
tvHWTxZYEcXLgAXFuUuaS3uF9gEiNQwzGTU1v0FqkqTBr4B8nW3HCN47XUu0t8Y0
e+lf4s4OxQawWD79J9/5d3Ry0vbV3Am1FtGJiJvOwRsIfVChDpYStTcHTCMqtvWb
V6L11BWkpzGXSW4Hv43qa+GSYOD2QU68Mb59oSk2OB+BtOLpJofmbGEGgvmwyCI9
MwIDAQAB
-----END PUBLIC KEY-----
`)

	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	publicKeyRSA := publicKey.(*rsa.PublicKey)

	claims, err := jwt.RSACheck(token, publicKeyRSA)
	if err != nil {
		log.Fatalf("RSACheck failed. %v", err)
	}

	log.Println("Expires: ", claims.Expires)

	if !claims.Valid(time.Unix(1516239099, 0)) {
		log.Fatal("Claim expired")
	}

	log.Println("RSACheck OK")
}
