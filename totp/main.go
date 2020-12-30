package main

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func main() {
	// key, err := totp.Generate(totp.GenerateOpts{
	// 	Period:      30,
	// 	Algorithm:   otp.AlgorithmSHA512,
	// 	Secret:      []byte("protrandat@gmail.comHENNGECHALLENGE003"),
	// 	Digits:      10,
	// 	SecretSize:  30,
	// 	Issuer:      "Hennge",
	// 	AccountName: "protrandat@gmail.com",
	// })

	secret := base32.StdEncoding.EncodeToString([]byte("protrandat@gmail.comHENNGECHALLENGE003"))
	code, err := totp.GenerateCodeCustom(
		secret,
		time.Now(),
		totp.ValidateOpts{
			Period:    30,
			Skew:      1,
			Digits:    10,
			Algorithm: otp.AlgorithmSHA512,
		})

	if err != nil {
		panic(err)
	}

	fmt.Println(code)

}

func sum(a, b int) int {
	return 0
}

func sumFloat(a, b float64) float64 {
	return 0.0
}
