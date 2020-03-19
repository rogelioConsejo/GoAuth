package main

func main()  {
	//carga automatica de servicios
	ConfiguracionInicial()

	//registro de nueva funcion en un servicio
	NuevaFuncion("Pagos","ActaulizarMoneda","id_moneda,nombre_moneda")

	//registro de servicios
	NuevoServicio("\"NUEVO\":{\"s\":\"0\",\"s\":\"3\"}")

	//exportacion de servicios
	ExportarConfiguracion()

}