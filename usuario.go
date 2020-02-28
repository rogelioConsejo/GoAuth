package goAuth

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type usuario struct {
	Email string
	password string
}

func NewUsuario(email string, password string) (usr *usuario, err error) {
	var emailValido bool
	var passwordValido bool
	usr = new(usuario)
	//TODO: Validar email y password
	emailValido, err = revisarEmail(email)
	if err == nil {
		passwordValido, err = revisarPassword(password)
	}

	if emailValido && passwordValido && err == nil{
		usr.Email = email
		//TODO: Cachar errores
		_ = usr.SetPassword(password)
	}

	return
}

func (u *usuario)SetPassword(newPassword string) (err error)  {
	//TODO: Revisar errores
	u.password, err = codificar(newPassword)
	return
}

func (u *usuario)CheckPassword(password string) (err error){
	passwordRevisado := []byte(password)
	passwordDeUsuario := []byte(u.password)
	err = bcrypt.CompareHashAndPassword(passwordDeUsuario, passwordRevisado)
	return
}


func codificar(stringInicial string) (stringCodificado string, err error){
	var inicial []byte
	var resultado []byte
	inicial = []byte(stringInicial)
	resultado, err = bcrypt.GenerateFromPassword(inicial, 12)
	stringCodificado = string(resultado)
	return
}

func revisarEmail(mail string)(esEmail bool, err error){
	esEmail, err = regexp.MatchString("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}", mail)
	return
}

func revisarPassword(password string) (esValido bool, err error) {
	//TODO: Implementar
	return
}