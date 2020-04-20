package auth

import (
	"encoding/json"
	"net/http"
	"strings"
)

//Función + parámetros
type Accion struct {
	nombreDelServicio      string
	verboHTTP              HttpVerb
	identificadorDeRecurso uri
	body                   requestBody
}

type Resultado struct {
}

type test_struct struct {
	Test string
}

func (a *Accion) Entity() (entity *accionEntity){
	entity = new(accionEntity)
	//_, nombreDelServicio, err := buscarAccionEnBaseDeDatos()
	//if err == nil{
		entity.nombreDelServicio = a.nombreDelServicio
		entity.verboHTTP = a.verboHTTP
		entity.identificadorDeRecurso = a.identificadorDeRecurso
		entity.body = a.body
	//}
	return
}

func CrearAccion()(a *Accion, err error){
	var nombreServicio string
	var verboHTTP HttpVerb
	var identificadorDeRecurso uri
	var body requestBody
	if nombreServicio != "" && verboHTTP != "" && identificadorDeRecurso != "" && body != ""{
		a = new(Accion)
		a.nombreDelServicio = nombreServicio
		a.verboHTTP = verboHTTP
		a.identificadorDeRecurso = identificadorDeRecurso
		a.body = body
		err = guardarAccion(a)
	}else if err == nil{
		a, err = CrearAccion()
	}
	return
}

func validarAccion(nombreAccion string)(a *Accion, err error){
	a, _, err = buscarAccion(nombreAccion)
	return
}
func destruirAccion(nombreServicio string)(err error){
	var id uint
	_, id, err = buscarAccion(nombreServicio)
	if err == nil{
		err = borrarAccion(id)
	}
	return
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
	decoder := json.NewDecoder(request.Body)
	reqToken := request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		// Error: Bearer token not in proper format
	}
	reqToken = strings.TrimSpace(splitToken[1])
	err, _ = obtenUsuarioPorToken(reqToken)

	var t test_struct
	err = decoder.Decode(&t)
	//err = obtenUsuarioPorToken(decoder)
	if err != nil {
		panic(err)
	}
	return
}
