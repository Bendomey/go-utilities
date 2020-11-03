package hashpassword

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes your password and returns a hash so that you can save that in your db
func HashPassword(password string) (string, error) {

	//pass to bcrypt to get byte slice
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	//return string
	return string(hashByte), err
}
