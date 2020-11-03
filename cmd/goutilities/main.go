package main

import (
	"fmt"

	"github.com/Bendomey/goutilities/pkg/hashpassword"
	"github.com/Bendomey/goutilities/pkg/validatehash"
)

func main() {
	// for genereating code
	// fmt.Println("my length ::", 6, "code generated :: ", generatecode.GenerateCode(6), "length :: ", len(generatecode.GenerateCode(6)))

	// for hashing password and compare
	a, _ := hashpassword.HashPassword("DomeyBenjamin")
	fmt.Println(validatehash.ValidateCipher("DomeyBenjamin", a))
}
