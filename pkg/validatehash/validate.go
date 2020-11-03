package validatehash

import "golang.org/x/crypto/bcrypt"

// ValidateCipher takes the password and hash and then compares the hashpassword to the password and see if they match
func ValidateCipher(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
