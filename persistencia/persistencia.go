package persistencia

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Conexion struct {
	DBdireccion string `json:"direccion"`
	DBpuerto    int    `json:"puerto"`    //No se usa, usa 3306
	DBusuario   string `json:"usuario"`
	DBPassword  string `json:"password"`
	DBnombre    string `json:"baseDeDatos"`
}

func conectarABaseDeDatos(c *Conexion) (baseDeDatos *sql.DB, err error) {
	var datosDeConexion string = fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true",
		c.DBusuario, c.DBPassword, c.DBdireccion,  c.DBnombre)
	baseDeDatos, err = sql.Open("mysql", datosDeConexion)

	if err != nil {
		mensajeDeError := fmt.Sprintf("no se pudo conectar a la base de datos (%s): %s\n",
			c.DBdireccion, err.Error())
		err = errors.New(mensajeDeError)
		baseDeDatos = nil
	}
	return
}

func cerrarConexion(baseDeDatos *sql.DB) (err error) {
	err = baseDeDatos.Close()
	return
}
