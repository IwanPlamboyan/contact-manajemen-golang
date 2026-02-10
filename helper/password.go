package helper

import "golang.org/x/crypto/bcrypt"

const bcryptCost = bcrypt.DefaultCost

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcryptCost,
	)
	return string(hashed), err
}

func ComparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}