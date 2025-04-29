package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pswd string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	return string(res), err
}

func CheckPasswordHash(pswd, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd)) == nil
}
