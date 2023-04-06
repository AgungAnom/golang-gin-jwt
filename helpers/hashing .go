package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) (res string, err error) {
	salt := 8
	arrByte := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(arrByte, salt)
	return string(hash),err
}

func ComparePass(h, p string) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}