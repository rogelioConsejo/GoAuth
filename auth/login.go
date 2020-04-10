package auth

import (
	"errors"
)

//TODO
//Intenta autenticar, si es exitoso, devuelve {codigo, nil}, en caso contrario devuelve {"", error}
//El codigo debe de guardarse en la base de datos para considerar que el login fue exitoso
func Login(email string, password string) (token string, err error) {
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
	if err != nil {
		println(usuario)
	}else {
		err=errors.New("Usuario o contraseña incorrectas")
	}
	//err = errors.New("no implementado")
	return
}

//TODO
//Elimina el codigo de la base de datos.
func Logout(token string) (err error) {
	err = errors.New("no implementado")
	return
}
