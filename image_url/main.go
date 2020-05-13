package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	key := "73757368690a"
	salt := "626164626f79640a"

	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}

	baseURL := "http://35.240.141.54"

	format := "jpg"
	id := "8583198232"
	options := "preset:view"

	path := fmt.Sprintf("/%s/plain/%s.%s", options, id, format)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	fmt.Printf("%s/%s%s", baseURL, signature, path)
}
