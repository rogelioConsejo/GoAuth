package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Usuario struct {
	email        string
	passwordHash string
}

func RevisarCredenciales(email string, password string) (usr *Usuario, err error) {
	usr, err = buscarUsuarioEnBaseDeDatos(email)
	if err == nil{
		err = usr.CheckPassword(password)
	}
	return
}

func NewUsuario(email string, password string) (usr *Usuario, err error) {
	usr = new(Usuario)

	err = usr.SetEmail(email)
	if err == nil{
		err = usr.SetPassword(password)
	}

	if err != nil {
		usr = nil
	}
	return
}

func (u *Usuario) SetPassword(newPassword string) (err error) {
	esValido, err := validarPassword(newPassword)
	if esValido && err == nil {
		u.passwordHash, err = codificar(newPassword)
	} else if !esValido {
		err = errors.New("password incorrecto")
	}
	return
}

func (u *Usuario) CheckPassword(password string) (err error) {
	passwordRevisado := []byte(password)
	passwordHashDeUsuario := []byte(u.passwordHash)
	err = bcrypt.CompareHashAndPassword(passwordHashDeUsuario, passwordRevisado)
	return
}

func (u *Usuario) SetEmail(email string) (err error) {
	esValido, err := validarEmail(email)
	if esValido && err == nil {
		u.email = email
	} else if err == nil {
		err = errors.New("email incorrecto")
	}
	return
}

func (u *Usuario) GetEmail() (email string){
	email = u.email
	return
}

func (u *Usuario) Registrar() (err error) {
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

func validarEmail(mail string) (esEmail bool, err error) {
	esEmail, err = regexp.MatchString("^[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}$", mail)
	return
}

func validarPassword(password string) (esValido bool, err error) {
	esValido, err = regexp.MatchString("^[A-Za-z0-9-+*/¡!#$%&?¿]{8,25}$", password)
	return
}

//TODO
func guardarUsuarioEnBaseDeDatos(u *Usuario) (err error){

	return
}

//TODO
func buscarUsuarioEnBaseDeDatos(email string)(u *Usuario, err error){
	return
}