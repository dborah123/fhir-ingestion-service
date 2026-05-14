package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("local-dev-secret-change-me")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":          "dev-client",
		"sourceSystem": "local",
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
	})

	signed, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	fmt.Println(signed)
}
