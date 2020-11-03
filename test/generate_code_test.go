package test

import (
	"testing"

	"github.com/Bendomey/goutilities/pkg/generatecode"
)

func TestGenerateCode(t *testing.T) {
	if len(generatecode.GenerateCode(6)) != 6 {
		t.Fatal("Couldn't genereate specified length")
	}
}
