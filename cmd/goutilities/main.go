package main

import (
	"fmt"

	"github.com/Bendomey/goutilities/pkg/generatecode"
)

func main() {
	fmt.Println("my length ::", 6, "code generated :: ", generatecode.GenerateCode(6), "length :: ", len(generatecode.GenerateCode(6)))
}
