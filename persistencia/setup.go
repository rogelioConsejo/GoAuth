package persistencia

import (
	"database/sql"
	"log"
)

//Crea una nueva estructura de base de datos vacía de acuerdo a una definición
func Instalar(conexion *Conexion, definicionDeBaseDeDatos *DefinicionDeBaseDeDatos) (err error) {
	log.Print("Iniciando instalación...")
	err = Configurar(conexion)

	var query string
	query, err = parsearBaseDeDatos(definicionDeBaseDeDatos)
	var db *sql.DB
	if err == nil {
		log.Printf("Instalando nueva instancia de API Gateway (Hecate): %s::%s@%s:%d\n",
			conexion.DBusuario, conexion.DBPassword, conexion.DBdireccion, conexion.DBpuerto)
		db, err = conectarABaseDeDatos(conexion)
	}

	//TODO: Revisar resultados
	//var res sql.Result
	if err == nil {
		_, err = db.Exec(query)
	}
	return
}



