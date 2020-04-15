package auth

import (
	"errors"
	"fmt"
)

//TODO
//Intenta autenticar, si es exitoso, devuelve {codigo, nil}, en caso contrario devuelve {"", error}
//El codigo debe de guardarse en la base de datos para considerar que el login fue exitoso
func Login(email string, password string) (t *token, err error) {
	/*usuario,id,err:=buscarUsuarioEnBaseDeDatos(email)
	if err!=nil {
		fmt.Print(usuario.passwordHash)
		passh,err:=codificar(password)
		if err!=nil {
			if usuario.passwordHash==passh {
				fmt.Print("login exitoso")
			}
		}
		fmt.Print(id)
	}*/
	/*passhas,err:=codificar(password)
	if err!=nil {
		usuario:=UsuarioEntity{Email:email,PasswordHash:passhas}
		row,err:=persistencia.BuscarEnBaseDeDatos(usuario,TABLA_USUARIOS)
		if err != nil {
			println(row)
		}else {
			err = errors.New("Usuario o contraseña incorrectos ")
		}
	} /*else {
		err = errors.New("")
	}*/
	usuario, err := RevisarCredenciales(email, password)
	if err == nil {
		fmt.Print(usuario)
		codigo:=generarToken(15)
		esUnico,err:=validarTokenUnico(codigo)

		if err == nil  && esUnico{
			t = new(token)
			t.codigo = codigo
			t.usr = usuario
			err = guardarToken(t)
			if err!=nil {
				err = errors.New("Fallo en el inicion de sesion intenta otra vez")
			}
		}
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
