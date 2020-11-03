package test

import (
	"testing"

	"github.com/Bendomey/goutilities/internal/calculate"
)

func TestCalculate(t *testing.T) {
	if calculate.Calculate(1, 2) != 3 {
		t.Fatal("Oops, calculate didn't work")
	}
}
