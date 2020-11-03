package main

import (
	"fmt"

	"github.com/Bendomey/goutilities/pkg/signjwt"
	"github.com/Bendomey/goutilities/pkg/validatetoken"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	// for genereating code
	// fmt.Println("my length ::", 6, "code generated :: ", generatecode.GenerateCode(6), "length :: ", len(generatecode.GenerateCode(6)))

	// for hashing password and compare
	// a, _ := hashpassword.HashPassword("DomeyBenjamin")
	// fmt.Println(validatehash.ValidateCipher("DomeyBenjamin", a))

	//for signing tokens
	claims := jwt.MapClaims{
		"user_id": 1,
	}
	token, _ := signjwt.SignJWT(claims, "helo")
	fmt.Println(validatetoken.ValidateJWTToken(token, "helo"))
}
