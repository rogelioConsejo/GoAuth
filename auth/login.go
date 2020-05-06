package auth

import (
	"errors"
	"fmt"
)

//TODO
//Intenta autenticar, si es exitoso, devuelve {Codigo, nil}, en caso contrario devuelve {"", error}
//El Codigo debe de guardarse en la base de datos para considerar que el login fue exitoso
func Login(email string, password string) (t *Token, err error) {
	usuario, err := RevisarCredenciales(email, password)
	if err == nil {
		t,err=crearToken(usuario)
	}else {
		err=errors.New("usuario o contraseña incorrectas")
	}
	return
}

//TODO
//Elimina el Codigo de la base de datos.
func Logout(token string) (err error) {
	err = destruirToken(token)
	if err==nil {
		fmt.Print("Sesion cerrada")
	}else{
		err=errors.New("Intenta de nuevo")
	}
	return err
}
