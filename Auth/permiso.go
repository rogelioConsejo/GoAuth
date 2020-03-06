package Auth

type permiso struct {
	nombre string
	puntosDeAcceso funciones
}

type funciones []funcionAPI

type funcionAPI struct {
	servicio string
	funcion string
	parametros []string
}

//TODO
func CrearPermiso(nombre string, puntosDeAcceso funciones) (p *permiso, err error){
	return
}

//TODO
func LeerPermisosRegistrados() (p []*permiso, err error){
	return
}

//TODO
func (p *permiso) agregarPuntoDeAcceso(puntoDeAcceso funcionAPI) (err error){
	return
}

//TODO
func (p *permiso) registrarPermiso()(err error){
	return
}

func (p *permiso) RevisarPermiso(u usuario)(tienePermiso bool){
	//TODO
	return 
}

func (p *permiso) OtorgarPermiso(u usuario)(err error){
	//TODO
	return
}