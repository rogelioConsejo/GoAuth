package persistencia

type Conexion struct {
	DBdireccion *string `json:"DBdireccion"`
	DBpuerto    *int    `json:"DBpuerto"`
	DBusuario   *string `json:"DBusuario"`
	DBPassword  *string `json:"DBPassword"`
}

func CambiarConfiguracion(conexion Conexion) {

}
