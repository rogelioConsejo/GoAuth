package persistencia

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Conexion struct {
	DBdireccion *string `json:"DBdireccion"`
	DBpuerto    *int    `json:"DBpuerto"`
	DBusuario   *string `json:"DBusuario"`
	DBPassword  *string `json:"DBpassword"`
}

func conectarABaseDeDatos(c *Conexion) (baseDeDatos *sql.DB, err error) {
	var datosDeConexion string = fmt.Sprintf("%s:%s@/%s:%d?parseTime=true",
		*c.DBusuario, *c.DBPassword, *c.DBdireccion, *c.DBpuerto)
	baseDeDatos, err = sql.Open("mysql", datosDeConexion)

	if err != nil {
		mensajeDeError := fmt.Sprintf("no se pudo conectar a la base de datos (%s): %s\n",
			*c.DBdireccion, err.Error())
		err = errors.New(mensajeDeError)
		baseDeDatos = nil
	}
	return
}

func cerrarConexion(baseDeDatos *sql.DB) (err error) {
	err = baseDeDatos.Close()
	return
}
