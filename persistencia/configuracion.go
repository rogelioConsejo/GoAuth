package persistencia

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

const maxPort = 65535
const configFilePath string = "db.conf"

//Configura la conexión a la base de datos
func Configurar(conexion *Conexion) (err error) {
	log.Printf("Configurando Base de Datos - %s@%s:%d\n",
		*conexion.DBusuario, *conexion.DBdireccion, *conexion.DBpuerto)
	err = validarConfiguracion(conexion)
	if err == nil {
		err = guardarConfiguracion(conexion)
	}
	return
}

//Realiza una validación simple de los datos de conexión a la base de datos y prueba la conexión
func validarConfiguracion(conexion *Conexion) (err error) {
	if *conexion.DBdireccion == "" {
		err = errors.New(fmt.Sprintf("dirección (%s) inválida", *conexion.DBdireccion))
	}
	if *conexion.DBpuerto <= 0 || *conexion.DBpuerto > maxPort {
		if err != nil {
			err = errors.New(fmt.Sprintf("puerto (%d) incorrecto, %s", *conexion.DBpuerto, err.Error()))
		} else {
			err = errors.New(fmt.Sprintf("puerto (%d) incorrecto", *conexion.DBpuerto))
		}
	}
	if *conexion.DBusuario == "" {
		if err != nil {
			err = errors.New(fmt.Sprintf("usuario incorrecto, %s", err.Error()))
		} else {
			err = errors.New("usuario incorrecto")
		}
	}

	if err != nil {
		err = errors.New(fmt.Sprintf("error en configuración de conexión a base de datos: %s", err.Error()))
	} else {
		var db *sql.DB
		db, err = conectarABaseDeDatos(conexion)
		if err == nil {
			err = db.Close()
		}
	}

	return
}

//Guarda la configuración para ejecuciones posteriores
func guardarConfiguracion(conexion *Conexion) (err error) {
	var file *os.File
	file, err = os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0755)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(*conexion)
	return err
}

//TODO: leer configuración existente y arrojar error si no existe ninguna
func getConfiguracion() (conexion *Conexion, err error) {
	var file *os.File
	file, err = os.OpenFile(configFilePath, os.O_RDONLY, 0755)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(conexion)
	return
}
