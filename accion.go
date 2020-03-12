package main

import "github.com/rogelioConsejo/Hecate/auth"

//Función + parámetros
type accion struct {
	permiso auth.Permiso
}

type resultado struct {

}

//TODO
func (a *accion) do(usr *auth.Usuario)(respuesta *resultado, err error){
	return
}

//TODO
func (a *accion) getIdentificador() (identificador string) {
	return
}

//TODO
func (r *resultado) getMensaje() (mensaje string){
	return
}