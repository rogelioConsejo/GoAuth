package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Token struct {
	Codigo string
	usr    *Usuario
}

func (t *Token) Entity() (entity *tokenEntity) {
	entity = new(tokenEntity)
	_, usrId, err := buscarUsuarioEnBaseDeDatos(t.usr.email)
	if err == nil {
		entity.Codigo = t.Codigo
		entity.Usuario = usrId
	}

	return
}

func crearToken(usuario *Usuario) (t *Token, err error) {
	var codigo string = fmt.Sprintf("%s-%s", generarToken(5), generarToken(20))
	var esUnico bool
	esUnico, err = validarTokenUnico(codigo)
	if err == nil && esUnico {
		t = new(Token)
		t.Codigo = codigo
		t.usr = usuario
		err = guardarToken(t)
	} else if err == nil {
		//Lo intenta hasta obtener un Codigo único, sucede una vez cada 12'000'000'000'000'000'000'000'000'000 años si se hace cada milisegundo
		t, err = crearToken(usuario)
	}
	return
}

func destruirToken(codigo string) (err error) {
	var id uint
	_, id, err = buscarToken(codigo)
	if err == nil {
		err = borrarToken(id)
	}
	return
}

func generarToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func validarTokenUnico(codigo string) (esUnico bool, err error) {
	var t *Token
	t, _, err = buscarToken(codigo)
	if err == nil {
		if t != nil {
			esUnico = false
		} else {
			esUnico = true
		}
	}
	if err.Error() == "sql: no rows in result set" {
		esUnico = true
		err = nil
	}

	return
}

func validarToken(codigo string)(t *Token, err error) {
	t, _, err = buscarToken(codigo)
	return
}