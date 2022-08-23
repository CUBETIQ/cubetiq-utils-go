package password

import "golang.org/x/crypto/bcrypt"

// encrypts user password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// check user password
func CheckPassword(password string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
