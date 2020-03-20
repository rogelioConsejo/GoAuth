package persistencia

import (
	"database/sql"
	"log"
)

//TODO
func Instalar(conexion *Conexion, definicionDeBaseDeDatos *BaseDeDatos) (err error) {
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
	//var res sql.Result
	if err == nil {
		_, err = db.Exec(query)
	}
	return
}



