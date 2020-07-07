package main

import (
	"errors"
	"github.com/rogelioConsejo/Hecate/auth"
	"log"
	"net/http"
	"strings"
)

//CÓDIGO PRINCIPAL DEL SERVIDOR
func handler(response http.ResponseWriter, request *http.Request) {

	/*var usuario *auth.Usuario
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
	tokent := request.Header.Get("tokent")
	fmt.Println(tokent)
	fmt.Println(request.Host)

	*/
	err := request.ParseForm()
	var token string
	var host string
	if request.Form != nil && err == nil {
		token = request.FormValue("token")
		host = request.FormValue("host")

		if token != " " || host != " " {
			url := strings.Split(request.Host, "/")
			log.Println(url)
			servicio, _, err := auth.BuscarServicio(url[0])
			log.Println(servicio)
			if err == nil && servicio != nil {
				redirectTLS(response, request, servicio.Direccion, "/"+url[1])
				log.Println(servicio.Direccion)
			} else {
				responsesErr(response, errors.New("no se encontro el servicio"), "Error ")
			}
		} else {
			responsesErr(response, errors.New(""), "No hay datos que porcesar")
		}
	} else {
		responsesErr(response, errors.New(""), "Error de formulario")
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

func loginHandler(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	var email string
	var password string

	if request.Form != nil && err == nil {

		email = request.FormValue("email")
		password = request.FormValue("pass")

		token, err := auth.Login(email, password)

		if err != nil {
			mensaje := "error en login ( u: " + email + " - p: " + password + " ): " + err.Error() + "\n"
			responsesErr(response, err, mensaje)
			return
		} else {
			responses(response, token.Codigo)
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

func logoutHandler(response http.ResponseWriter, request *http.Request) {

	var token string
	var err error

	token = request.FormValue("token")

	if token != " " {
		err = auth.Logout(token)
		if err != nil {
			responsesErr(response, err, "Error de logout")
		}
	} else {
		responsesErr(response, err, "Error no hay token")
	}
}

func serviciosHandler(response http.ResponseWriter, request *http.Request) {

	var accion string
	var direccion string
	var nombre string
	//var id_servicio string
	var err error

	accion = request.FormValue("accion")
	direccion = request.FormValue("direccion")
	nombre = request.FormValue("nombre")

	if accion != "" || direccion != "" || nombre != "" {

		servicio := auth.Servico{Nombre: nombre, Direccion: direccion}

		switch accion {

		case "registrar":
			_, id, err := auth.BuscarServicio(nombre)
			if id == 0 {
				if err != nil {
					_, err := auth.RegistrarServicio(servicio)
					if err != nil {
						responsesErr(response, err, "Error al registrar")
					} else {
						responses(response, "Registro exitoso")
					}
				}
			} else {
				responses(response, "Datos ya existentes")
			}
			break
		case "buscar":
			_, id, err := auth.BuscarServicio(nombre)
			if err == nil {
				responses(response, string(id))
			} else {
				responsesErr(response, err, "")
			}
			break
		case "actualizar":
			//err := auth.ActualizarServicio(id_servicio)
			break
		case "eliminar":
			err := auth.EliminarServicio(nombre)
			if err == nil {
				responses(response, "eliminacion correcta")
			} else {
				responsesErr(response, err, "Error en eliminacion")
			}
			break
		default:
			responses(response, "Accion desconocida")
			break
		}
	} else {
		responsesErr(response, err, "Error no hay datos que procesar")
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request, servidor string, accion string) {
	http.Redirect(w, r, servidor+accion+r.RequestURI, http.StatusMovedPermanently)
}

func responses(response http.ResponseWriter, mensaje string) {
	var err error

	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	response.Header().Set("Access-Control-Allow-Methods", "POST")
	response.Header().Set("Content-Type", "text/plain")

	_, err = response.Write([]byte(mensaje))

	if err == nil {
		log.Printf("%v\n", mensaje)
	} else {
		log.Println("Error de escritura")
	}
}

func responsesErr(response http.ResponseWriter, err error, mensaje string) {
	http.Error(response, err.Error(), http.StatusUnauthorized)
	log.Println(mensaje, ' ', err.Error())
}
