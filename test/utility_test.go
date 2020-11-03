package test

import (
	"testing"

	"github.com/Bendomey/goutilities/pkg/calculate"
	"github.com/Bendomey/goutilities/pkg/generatecode"
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
