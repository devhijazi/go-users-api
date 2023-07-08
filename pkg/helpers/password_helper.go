package helpers

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hash)
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
