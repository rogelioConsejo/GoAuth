package persistencia

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

const maxPort = 65535
const configFilePath string = "C:\\Users\\Roberto\\go\\src\\Hecate\\persistencia\\db.conf"

//Configura la conexión a la base de datos
func Configurar(conexion *ConfiguracionDeConexion) (err error) {
	log.Printf("Configurando Base de Datos - %s:[password]@%s/%s\n",
		*conexion.DBusuario, *conexion.DBdireccion, *conexion.DBnombre)
	err = validarConfiguracion(conexion)
	if err == nil {
		err = guardarConfiguracion(conexion)
	}
	return
}

//Realiza una validación simple de los datos de conexión a la base de datos y prueba la conexión
func validarConfiguracion(conexion *ConfiguracionDeConexion) (err error) {
	if *conexion.DBusuario == "" {
		err = errors.New("usuario incorrecto")
	}
	if *conexion.DBnombre == "" {
		err = errors.New("nombre de base de datos incorrecto")
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("error en configuración de conexión a base de datos: %s", err.Error()))
	} else {
		err = probarConexion(err, conexion)
	}
	return
}

//Guarda la configuración para ejecuciones posteriores
func guardarConfiguracion(conexion *ConfiguracionDeConexion) (err error) {
	var file *os.File
	file, err = os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0755)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(*conexion)
	return err
}

//Lee la configuración de conexión a base de datos existente
func getConfiguracion() (conexion *ConfiguracionDeConexion, err error) {
	conexion = new(ConfiguracionDeConexion)
	var file *os.File
	file, err = os.OpenFile(configFilePath, os.O_RDONLY, 0755)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(conexion)
	referencia := new(ConfiguracionDeConexion)
	if *conexion == *referencia {
		if err != nil {
			errString := fmt.Sprintf("no se recuperó ninguna configuración existente: %s\n", err.Error())
			err = errors.New(errString)
		} else {
			err = errors.New("no parece haber una configuración definida")
		}
	}
	return
}
