package validatetoken

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// ValidateJWTToken takes in the signedtoken and secret and then returns your data
func ValidateJWTToken(tokenString string, secret string) (*jwt.Token, error) {
	data, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return data, err
}
