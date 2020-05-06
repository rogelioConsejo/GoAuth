package main

import (
	"github.com/rogelioConsejo/Hecate/auth"
	"log"
	"net/http"
)

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
		if resultado != nil && resultado.GetMensaje() != nil {
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

func loginHandelr(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	var email string
	var password string

	if request.Form != nil && err == nil {

		email = request.FormValue("email")
		password = request.FormValue("pass")

		token, err := auth.Login(email, password)

		if err != nil {
			http.Error(response, err.Error(), http.StatusUnauthorized)
			log.Printf("error en login ( u: %s - p: %s ): %s\n", email, password, err.Error())
			return
		} else {
			response.Header().Set("Access-Control-Allow-Origin", "*")
			response.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			response.Header().Set("Access-Control-Allow-Methods", "POST")
			response.Header().Set("Content-Type", "text/plain")
			_, err = response.Write([]byte(token.Codigo))
			log.Printf("Login exitoso ( u:%s )\n", email)
		}

	} else {
		errorString := ""
		http.Error(response, "error al leer datos de login", http.StatusInternalServerError)
		if err != nil {
			errorString = ": " + err.Error()
		}
		log.Printf("error al parsear datos de login o datos vacíos%s\n", errorString)
	}
}
