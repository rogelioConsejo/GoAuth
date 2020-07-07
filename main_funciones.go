package main

import (
	"encoding/json"
	"flag"
	"github.com/rogelioConsejo/Hecate/persistencia"
	"log"
	"os"
)

//Definición de banderas
func leerBanderas() (p parametros) {
	p.Configuracion = new(ConfiguracionDeServidor)
	p.ConfiguracionDeBD = new(persistencia.ConfiguracionDeConexion)
	p.EsInstalacion = flag.Bool("nuevo", false,
		"Indica que se quiere instalar por primera vez, borrando todas las tablas existentes")
	p.EsConfiguracion = flag.Bool("config", false,
		"Indica que se quiere configurar el programa")

	p.ConfiguracionDeBD.DBnombre = flag.String("db", "hecate", "El nombre de la base de datos")
	p.ConfiguracionDeBD.DBdireccion = flag.String("dbdir", "127.0.0.1", "La dirección de la base de datos")
	p.ConfiguracionDeBD.DBpuerto = flag.Int("dbport", 3306, "El puerto de la base de datos")
	p.ConfiguracionDeBD.DBusuario = flag.String("dbusr", "root",
		"El nombre de usuario a usar para la ConfiguracionDeConexion a base de datos")
	p.ConfiguracionDeBD.DBPassword = flag.String("dbpass", "",
		"El password a usar para la ConfiguracionDeConexion a base de datos")

	p.Configuracion.DireccionDeServidor = flag.String("d", "localhost",
		"La direccion en donde será accesible el servidor, se debe definir también un puerto")
	p.Configuracion.PuertoDeServidor = flag.Uint("p", 8080,
		"El puerto desde donde será accesible el servidor, se debe definir también una dirección")

	flag.Parse()
	log.Printf("nombre de base de datos: %s\n", *p.ConfiguracionDeBD.DBnombre)

	return
}


//Obtiene la definición de la base de datos desde definicionBD.xconf
func generarDefinicionDeBaseDeDatos() (definicion *persistencia.DefinicionDeBaseDeDatos, err error) {
	file, err := os.Open("definicionBD.xconf")
	definicion = new(persistencia.DefinicionDeBaseDeDatos)

	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(definicion)
	}

	return
}
