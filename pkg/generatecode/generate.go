package generatecode

import (
	"math/rand"
	"strconv"
	"time"
)

// GenerateCode takes in the length and generate code. Parameter is required
func GenerateCode(length int) string {
	// get the max from the length
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := "9"
	for i := 1; i < length; i++ {
		max += "9"
	}
	i, _ := strconv.Atoi(max)
	return strconv.Itoa(rand.Intn(i-min+1) + min)
}
