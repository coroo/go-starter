package utils

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func StringToInt(req string) int {
	res, _ := strconv.Atoi(req)
	return res
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
