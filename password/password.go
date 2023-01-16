package password

import "golang.org/x/crypto/bcrypt"

// encrypts raw password
func HashPassword(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	return string(bytes), err
}

// compare raw password with hash password
func CheckPassword(rawPassword string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))

	return err == nil
}
