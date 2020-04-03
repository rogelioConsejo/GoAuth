package auth

import "errors"

//TODO
//Intenta autenticar, si es exitoso, devuelve {token, nil}, en caso contrario devuelve {"", error}
//El token debe de guardarse en la base de datos para considerar que el login fue exitoso
func Login(email string, password string) (token string, err error) {
	err = errors.New("no implementado")
	return
}

//TODO
//Elimina el token de la base de datos.
func Logout(token string)(err error) {
	err = errors.New("no implementado")
	return
}

