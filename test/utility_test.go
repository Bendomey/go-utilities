package test

import (
	"testing"

	"github.com/Bendomey/goutilities/pkg/calculate"
	"github.com/Bendomey/goutilities/pkg/generatecode"
	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"github.com/Bendomey/goutilities/pkg/validatehash"
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
