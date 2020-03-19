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
		conexion.DBusuario, conexion.DBdireccion, conexion.DBpuerto)
	err = validarConfiguracion(conexion)
	if err == nil {
		err = guardarConfiguracion(conexion)
	}
	return
}

//Realiza una validación simple de los datos de conexión a la base de datos y prueba la conexión
func validarConfiguracion(conexion *Conexion) (err error) {
	if conexion.DBusuario == "" {

		err = errors.New("usuario incorrecto")

	}
	if conexion.DBnombre == "" {
		err = errors.New("nombre de base de datos incorrecto")
	}

	if err != nil {
		err = errors.New(fmt.Sprintf("error en configuración de conexión a base de datos: %s", err.Error()))
	} else {
		var db *sql.DB
		db, err = conectarABaseDeDatos(conexion)
		if err == nil {
			err = db.Ping()
		}
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

//Lee la configuración de conexión a base de datos existente
func getConfiguracion() (conexion *Conexion, err error) {
	conexion = new(Conexion)
	var file *os.File
	file, err = os.OpenFile(configFilePath, os.O_RDONLY, 0755)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(conexion)
	referencia := new(Conexion)
	if *conexion == *referencia {
		if err != nil{
			errString := fmt.Sprintf("no se recuperó ninguna configuración existente: %s\n", err.Error())
			err = errors.New(errString)
		} else {
			err = errors.New("no parece haber una configuración definida")
		}
	}
	return
}
