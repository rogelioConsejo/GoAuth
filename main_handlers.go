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
	var accionARealizar *accion
	var err error

	accionARealizar, usuario, err = parsearPeticion(request)

	if tienePermiso && err == nil {
		log.Printf("Petición (%s): %s\n", usuario.GetEmail(), accionARealizar.getIdentificador())
		var resultado *resultado
		resultado, err = accionARealizar.do(usuario)
		if resultado != nil {
			log.Printf("Resultado (%s): %s -> %s\n",
				usuario.GetEmail(), accionARealizar.getIdentificador(), resultado.getMensaje())
		}
	} else if !tienePermiso && err == nil {
		log.Printf("ALERTA: usuario %s intentó realizar una acción sin permiso: %s\n", usuario.GetEmail(),
			accionARealizar.getIdentificador())
	}

	if err != nil {
		log.Printf("error en API Gateway: %s\n", err.Error())
	}
}

func usrHandler(response http.ResponseWriter, request *http.Request) {
	var err error
	var accion *accion
	var resultado *resultado
	var usuario *auth.Usuario

	accion, usuario, err = parsearPeticion(request)

	resultado, err = accion.do(usuario)
	if err == nil {
		log.Print(resultado.getMensaje())
	} else {
		log.Printf("error al realizar acción de usuario: %s\n", err.Error())
	}
}
