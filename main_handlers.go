package main

import (
	"encoding/json"
	"fmt"
	"github.com/rogelioConsejo/Hecate/auth"
	"log"
	"net/http"
)

type Response struct {
	status    	bool
	messages 	string
	log			string
}

//CÓDIGO PRINCIPAL DEL SERVIDOR
func handler(response http.ResponseWriter, request *http.Request) {
	var usuario *auth.Usuario
	var tienePermiso bool
	var accionARealizar *auth.Accion
	var err error

	accionARealizar, usuario, err = auth.ParsearPeticion(request)

	if tienePermiso && err == nil {
		log.Printf("Petición (%s): %s\n", usuario.GetEmail(), accionARealizar.GetNombre())
		var resultado *auth.Resultado
		resultado, err = accionARealizar.Do(usuario)
		if resultado != nil && resultado.GetMensaje() != nil{
			log.Printf("Resultado (%s): %s -> %s\n",
				usuario.GetEmail(), accionARealizar.GetNombre(), *resultado.GetMensaje())
		}
	} else if !tienePermiso && err == nil {
		log.Printf("ALERTA: usuario %s intentó realizar una acción sin permiso: %s\n", usuario.GetEmail(),
			accionARealizar.GetNombre())
	}

	if err != nil {
		log.Printf("error en API Gateway: %s\n", err.Error())
	}
}

func usrHandler(response http.ResponseWriter, request *http.Request) {
	var err error
	var accion *auth.Accion
	var resultado *auth.Resultado
	var usuario *auth.Usuario

	accion, usuario, err = auth.ParsearPeticion(request)

	resultado, err = accion.Do(usuario)
	if err == nil {
		log.Print(resultado.GetMensaje())
	} else {
		log.Printf("error al realizar acción de usuario: %s\n", err.Error())
	}
}


func loginHandelr(response http.ResponseWriter, request *http.Request)  {
	request.ParseForm()
	var resp Response
	var email string
	var password string

	if request.Form != nil {

	email=request.FormValue("email")
	password=request.FormValue("pass")



	log.Print(request.FormValue("email"))
	log.Print(request.FormValue("pass"))

	token,_,err:=auth.Login(email,password)

	fmt.Print(token)

	if err!=nil {
		resp=Response{status: false,messages: "Error ",log: ""}
	}else {
		resp=Response{status: true,messages: "",log: token}
	}
	
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
	}else{

		resp :=Response{status: false,messages: "Sin informacion que porcesar",log: ""}

		js, err := json.Marshal(resp)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		response.Write(js)
	}
}