package main

import (
	"flag"
	"github.com/rogelioConsejo/Hecate/configuracion"
	"net/http"
)

type conexion struct {
	direccionDB *string
	puertoDB *int
	usuarioDB *string
	
}

func main() {
	//TODO: banderas
	var esInstalacion *bool = flag.Bool("nuevo", false,
		"Indica que se quiere instalar por primera vez, borrando todas las tablas existentes")
	var esConfiguracion *bool = flag.Bool("config", false,
		"Indica que se quiere configurar el programa")
	var c conexion
	c.direccionDB = flag.String("db","","La dirección de la base de datos")
	c.puertoDB = flag.Int("puertoDB",3306, "El puerto de la base de datos")

	flag.Parse()

	if *esInstalacion {
		configuracion.Instalar()
	} else if *esConfiguracion{
		configuracion.CambiarConfiguracion()
	}

	//TODO: Set-up base de datos SQL con banderas - usando un archivo de configuración

	//TODO: Correr servidor con dirección configurable por banderas

	//http.HandleFunc("/", handler)
	//err := http.ListenAndServe("localhost:8080", nil)
	//if err != nil {
	//	//TODO: Loggear error de montado de servidor
	//}
}

func handler(response http.ResponseWriter, request *http.Request) {

}
