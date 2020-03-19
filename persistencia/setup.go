package persistencia

import (
	"log"
)

//TODO
func Instalar(conexion *Conexion, definicionDeBaseDeDatos *BaseDeDatos) (err error) {
	err = guardarConfiguracion(conexion)

	if err == nil {
		log.Printf("Instalando nueva instancia de API Gateway (Hecate): %s::%s@%s:%d\n",
			conexion.DBusuario, conexion.DBPassword, conexion.DBdireccion, conexion.DBpuerto)
		//TODO: Instalar base de datos
	}
	return
}



