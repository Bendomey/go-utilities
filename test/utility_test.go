package test

import (
	"testing"

	"github.com/Bendomey/goutilities/pkg/calculate"
	"github.com/Bendomey/goutilities/pkg/generatecode"
	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"github.com/Bendomey/goutilities/pkg/signjwt"
	"github.com/Bendomey/goutilities/pkg/validatehash"
	"github.com/Bendomey/goutilities/pkg/validatetoken"
	"github.com/dgrijalva/jwt-go"
)

func TestCalculate(t *testing.T) {
	if calculate.Calculate(1, 2) != 3 {
		t.Fatal("Oops, calculate didn't work")
	}
}

func TestGenerateCode(t *testing.T) {
	if len(generatecode.GenerateCode(6)) != 6 {
		t.Fatal("Couldn't genereate specified length")
	}
}

func TestHashPassword(t *testing.T) {
	_, err := hashpassword.HashPassword("DOmeyBenjamin")
	if err != nil {
		t.Fatalf("SOmehing happened whiles trying to hash password :: %v", err)
	}
}

func TestComparePasswords(t *testing.T) {
	password := "Domey Benjamin"
	hash, _ := hashpassword.HashPassword(password)
	isSame := validatehash.ValidateCipher(password, hash)
	if isSame == false {
		t.Fatal("Oops, password could not be compared")
	}
}

func TestSignToken(t *testing.T) {
	secretForJWT := "hello"
	claims := jwt.MapClaims{}
	claims["user_id"] = 1
	_, err := signjwt.SignJWT(claims, secretForJWT)
	if err != nil {
		t.Fatalf("Oopsn jwt couldn't sign token :: %v", err)
	}
}

func TestValidateToken(t *testing.T) {
	secretForJWT := "hello"
	claims := jwt.MapClaims{}
	claims["user_id"] = 1
	token, err := signjwt.SignJWT(claims, secretForJWT)
	if err != nil {
		t.Fatalf("Oopsn jwt couldn't sign token :: %v", err)
	}

	_, errValidate := validatetoken.ValidateJWTToken(token, secretForJWT)
	if errValidate != nil {
		t.Fatalf("Oopsn jwt couldn't validate token :: %v", errValidate)

	}
}
