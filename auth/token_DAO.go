package auth

import "errors"

func guardarToken (t *token) (err error) {
	err = errors.New("no implementado")

	return
}

func borrarToken (t *token) (err error) {
	err = errors.New("no implementado")

	return
}

func buscarToken(codigo string)(t *token, err error)  {
	err = errors.New("no implementado")
	t = nil

	return
}