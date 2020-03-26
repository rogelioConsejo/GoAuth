package auth

import "errors"

func Login(email string, password string) (token string, err error) {
	err = errors.New("no implementado")
	return
}

func Logout(token string)(err error) {
	err = errors.New("no implementado")
	return
}

