package main

import (
	"github.com/rogelioConsejo/Hecate/auth"
	"net/http"
)

//Función + parámetros
type accion struct {
}

type resultado struct {
}

type funcion func(args map[string]Argumento)

type Argumento struct {
	tipo  *TipoArgumento
	valor string
}

type TipoArgumento struct {

}

//TODO
func (a *accion) do(usr *auth.Usuario) (respuesta *resultado, err error) {
	return
}

//TODO
func (a *accion) getIdentificador() (identificadorDeAccion string) {
	return
}

//TODO
func (r *resultado) getMensaje() (mensaje string) {
	return
}

//TODO: Implementar
func parsearPeticion(request *http.Request) (accion *accion, usuario *auth.Usuario, err error) {

	return
}
