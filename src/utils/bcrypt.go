package utils

import "golang.org/x/crypto/bcrypt"

func HashMake(password string) (hashed string, err error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(res), err
}

func HashValidate(password, hashed string) (isPasswordValidated bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
