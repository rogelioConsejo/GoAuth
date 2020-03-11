package auth

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
func (p *permiso) Registrar()(err error){
	return
}

func (p *permiso) Revisar(u usuario)(tienePermiso bool){
	//TODO
	return 
}

func (p *permiso) Otorgar(u usuario)(err error){
	//TODO
	return
}