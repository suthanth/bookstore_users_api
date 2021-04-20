package crypto_utils

import (
	"github.com/suthanth/bookstore_users_api/utils/rest_errors"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, *rest_errors.RestErr) {
	if password == "" {
		return "", rest_errors.NewBadRequest("Password cannot be null")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", rest_errors.NewBadRequest(err.Error())
	}
	return string(hash), nil
}

func ComparePassword(hashedPassword, password string) bool {
	if hashedPassword == "" || password == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
