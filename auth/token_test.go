package auth

import (
	"fmt"
	"testing"
)

func TestGenerarToken(t *testing.T) {
	var i uint8
	for i = 1; i < 10 ; i++  {
		fmt.Printf("%s-%s\n",generarToken(5),generarToken(20))
	}
	generarToken(20)
}

func TestToken(t  *testing.T) {
	var err error
	var usuario *Usuario
	var token *Token

	usuario, _, err = buscarUsuarioEnBaseDeDatos("rogelio.consejo@gmail.com")
	if err == nil {
		token, err = crearToken(usuario)
	}
	fmt.Printf("%+v\n", token)

	var codigo string
	if err == nil {
		codigo = token.Codigo
		fmt.Printf("código: %s\n", codigo)
		token, err = validarToken(codigo)
	}
	fmt.Printf("%+v\n", token)
	if err == nil {
		fmt.Printf("usuario asociado: %s\n", token.usr.email)
	}

	var id uint
	_, id, err = buscarToken(codigo)
	if err == nil {
		err = borrarToken(id)
	}

	if err != nil {
		t.Errorf("prueba fallada: %s\n", err.Error())
	}

}