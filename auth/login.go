package auth

import (
	"errors"
	"fmt"
)

//TODO
//Intenta autenticar, si es exitoso, devuelve {codigo, nil}, en caso contrario devuelve {"", error}
//El codigo debe de guardarse en la base de datos para considerar que el login fue exitoso
func Login(email string, password string) (r string,t *token, err error) {
	usuario, err := RevisarCredenciales(email, password)
	if err == nil {
		//fmt.Print(usuario)
		t,err=crearToken(usuario)
		r=t.codigo
	}else {
		err=errors.New("Usuario o contraseña incorrectas")
	}
	//err = errors.New("no implementado")
	return
}

//TODO
//Elimina el codigo de la base de datos.
func Logout(token string) (err error) {
	//err = errors.New("no implementado")
	err = destruirToken(token)
	if err==nil {
		fmt.Print("Sesion cerrada")
	}else{
		err=errors.New("Intenta de nuevo")
	}
	return err
}
