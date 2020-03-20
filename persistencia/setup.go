package persistencia

import (
	"database/sql"
	"log"
	"strings"
)

//Crea una nueva estructura de base de datos vacía de acuerdo a una definición
func Instalar(conexion *ConfiguracionDeConexion, definicionDeBaseDeDatos *DefinicionDeBaseDeDatos) (err error) {
	log.Print("Iniciando instalación...")

	//Convertir DefinicionDeBaseDeDatos en query
	var queriesString string
	queriesString, err = parsearBaseDeDatos(definicionDeBaseDeDatos)

	//Configurar conexión
	err = Configurar(conexion)
	var db *sql.DB
	if err == nil {
		log.Printf("Instalando nueva instancia de API Gateway (Hecate): %s::%s@%s\n",
			*conexion.DBusuario, *conexion.DBPassword, *conexion.DBdireccion)
		db, err = conectarABaseDeDatos(conexion)
	}

	//Ejecutar queries
	var res sql.Result
	if err == nil {
		var queries []string
		queries = strings.Split(queriesString, "\n")

		for _, query := range queries {
			if query != ""{
				if err == nil {
					res, err = db.Exec(query)
				}
				if err == nil {
					err = loggearResultadoDeQuery(err, res, query)
				}
			}

		}

	}

	return
}

//Loggea el resultado de un query Exec
func loggearResultadoDeQuery(err error, res sql.Result, query string) error {
	var numRows int64
	numRows, err = res.RowsAffected()
	log.Printf("query: %s\n", query)
	log.Printf("rows affected: %d\n", numRows)
	return err
}
