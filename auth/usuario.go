package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/rogelioConsejo/Hecate/persistencia"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	TABLA_USUARIOS = "usuarios"
)

type Usuario struct {
	email        string
	permaToken   string //Se usa para comunicación intraservicios
	passwordHash string
}

type UsuarioEntity struct {
	Id           uint
	Email        string
	PermaToken   string
	PasswordHash string
}

func (u UsuarioEntity) GetId() uint {
	return u.Id
}

func RevisarCredenciales(email string, password string) (usr *Usuario, err error) {
	usr, _, err = buscarUsuarioEnBaseDeDatos(email)
	if err == nil {
		err = usr.CheckPassword(password)
	}
	return
}

func NewUsuario(email string, password string) (usr *Usuario, err error) {
	usr = new(Usuario)

	err = usr.SetEmail(email)
	if err == nil {
		err = usr.SetPassword(password)
	}
	if err == nil {
		err = usr.crearPermaToken()
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

func (u *Usuario) GetEmail() (email string) {
	email = u.email
	return
}

func (u *Usuario) Registrar() (err error) {
	err = guardarUsuarioEnBaseDeDatos(u)
	return
}

func (u *Usuario) crearPermaToken() (err error) {
	var permaToken string
	if u == nil {
		err = errors.New("no se puede crear permatoken en un usuario no definido")
	}
	if err == nil {
		var esUnico bool = false
		for !esUnico {
			permaToken = fmt.Sprintf("%s::%s{%s}", generarToken(4), generarToken(5), generarToken(25))
			esUnico, err = revisarPermaTokenUnico(permaToken)
		}
	}
	if err == nil {
		u.permaToken = permaToken
	}

	return
}

func (u *Usuario) Entity() *UsuarioEntity {
	entity := new(UsuarioEntity)
	entity.Email = u.email
	entity.PermaToken = u.permaToken
	entity.PasswordHash = u.passwordHash
	return entity
}

//TODO
func revisarPermaTokenUnico(permaToken string) (bool, error) {
	//return false, errors.New("no implementado")
	return true, nil
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
	esEmail, err = regexp.MatchString("^[a-zA-Z0-9\\.\\-]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}$", mail)
	return
}

func validarPassword(password string) (esValido bool, err error) {
	esValido, err = regexp.MatchString("^[A-Za-z0-9-+*/¡!#$%&?¿]{8,25}$", password)
	return
}

func guardarUsuarioEnBaseDeDatos(u *Usuario) (err error) {
	_, err = persistencia.RegistrarEnBaseDeDatos(*u.Entity(), TABLA_USUARIOS)
	return
}

func buscarUsuarioEnBaseDeDatos(email string) (u *Usuario, id uint, err error) {
	u = new(Usuario)
	e := new(UsuarioEntity)
	u.email = email
	rows, err := persistencia.BuscarUnoEnBaseDeDatos(*u.Entity(), TABLA_USUARIOS)
	if err == nil {
		err = rows.Scan(&e.Id,&e.Email,&e.PermaToken, &e.PasswordHash)
	}
	if err == nil {
		id = e.Id
		u.email = e.Email
		u.passwordHash = e.PasswordHash
		u.permaToken = e.PermaToken
	} else {
		u = nil
	}
	return
}

func leerUsuarioEnBaseDeDatos(id uint)(u *Usuario, err error) {
	e := new(UsuarioEntity)

	var row *sql.Row
	row, err = persistencia.LeerEnBaseDeDatos(id, *e, TABLA_USUARIOS)
	if err == nil {
		err = row.Scan(&e.Id,&e.Email,&e.PermaToken,&e.PasswordHash)
	}

	u = new(Usuario)
	if err == nil {
		u.email = e.Email
		u.permaToken = e.PermaToken
		u.passwordHash = e.PasswordHash
	}
	return
}
