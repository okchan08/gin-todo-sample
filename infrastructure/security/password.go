package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, 12)

	return string(hash), err
}

func ComparePasswordAndHash(password string, hash string) bool {
	bytePass := []byte(password)
	byteHash := []byte(hash)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)

	return err == nil
}
