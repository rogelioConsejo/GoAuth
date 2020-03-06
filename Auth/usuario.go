package Auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type usuario struct {
	email    string
	password string
}

func NewUsuario(email string, password string) (usr *usuario, err error) {
	usr = new(usuario)

	err = usr.SetEmail(email)
	if err == nil{
		err = usr.SetPassword(password)
	}

	if err != nil {
		usr = nil
	}
	return
}

func (u *usuario) SetPassword(newPassword string) (err error) {
	esValido, err := revisarPassword(newPassword)
	if esValido && err == nil {
		u.password, err = codificar(newPassword)
	} else if !esValido {
		err = errors.New("password incorrecto")
	}
	return
}

func (u *usuario) CheckPassword(password string) (err error) {
	passwordRevisado := []byte(password)
	passwordDeUsuario := []byte(u.password)
	err = bcrypt.CompareHashAndPassword(passwordDeUsuario, passwordRevisado)
	return
}

func (u *usuario) SetEmail(email string) (err error) {
	esValido, err := revisarEmail(email)
	if esValido && err == nil {
		u.email = email
	} else if err == nil {
		err = errors.New("email incorrecto")
	}
	return
}

func (u *usuario) GetEmail() (email string){
	email = u.email
	return
}

//TODO
func (u *usuario) Registrar() (err error) {
	err = guardarUsuarioEnBaseDeDatos(u)
	return
}

func codificar(stringInicial string) (stringCodificado string, err error) {
	var inicial []byte
	var resultado []byte
	inicial = []byte(stringInicial)
	resultado, err = bcrypt.GenerateFromPassword(inicial, 12)
	stringCodificado = string(resultado)
	return
}

func revisarEmail(mail string) (esEmail bool, err error) {
	esEmail, err = regexp.MatchString("^[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}$", mail)
	return
}

func revisarPassword(password string) (esValido bool, err error) {
	esValido, err = regexp.MatchString("^[A-Za-z0-9-+*/¡!#$%&?¿]{8,25}$", password)
	return
}

//TODO
func guardarUsuarioEnBaseDeDatos(u *usuario) (err error){
	return
}
