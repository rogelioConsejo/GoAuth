package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type token struct {
	id    uint32
	token string
	usr   *Usuario
}

func crearToken(usuario *Usuario)(t *token, err error) {
	var codigo string = fmt.Sprintf("%s-%s",generarToken(5),generarToken(20))
	var esUnico bool
	esUnico, err = validarTokenUnico(codigo)
	if err == nil && esUnico {
		t.token = codigo
		t.usr = usuario
		err = guardarToken(t)
	} else if err == nil {
		//Lo intenta hasta obtener un token único, sucede una vez cada 12'000'000'000'000'000'000'000'000'000 años si se hace cada milisegundo
		t, err = crearToken(usuario)
	}
	return
}



func destruirToken(codigo string) (err error) {
	var t *token
	t, err = buscarToken(codigo)
	if err == nil {
		err = borrarToken(t)
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
	var t *token
	t, err = buscarToken(codigo)
	if err == nil {
		if t != nil {
			esUnico = false
		} else {
			esUnico = true
		}
	}

	return
}