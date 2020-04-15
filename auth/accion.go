package auth

import (
	"net/http"
)

//Función + parámetros
type Accion struct {
	nombre                 string
	nombreDelServicio      string
	verboHTTP              HttpVerb
	identificadorDeRecurso uri
	body                   requestBody
}

type Resultado struct {
}

//TODO
func (a *Accion) Do(usr *Usuario) (respuesta *Resultado, err error) {
	return
}

//TODO
func (a *Accion) GetNombre() (identificadorDeAccion string) {
	return
}

//TODO
//Mensaje de Resultado apto para mostrarse al usuario, si no hay es nil
func (r *Resultado) GetMensaje() (mensaje *string) {
	mensaje = nil
	return
}

//TODO
//Mensaje para log, apto sólo para uso interno
func (r *Resultado) GetLog() (log string) {
	return
}

//TODO: Implementar
func ParsearPeticion(request *http.Request) (accion *Accion, usuario *Usuario, err error) {
	return
}
