package signjwt

import (
	"github.com/dgrijalva/jwt-go"
)

// SignJWT takes in jwt claims and secret and then returns a signed token or an error
func SignJWT(claims jwt.MapClaims, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(secret))
	return token, err
}
