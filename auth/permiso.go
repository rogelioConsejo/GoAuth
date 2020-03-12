package auth

type Permiso struct {
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
func CrearPermiso(nombre string, puntosDeAcceso funciones) (p *Permiso, err error){
	return
}

//TODO
func LeerPermisosRegistrados() (p []*Permiso, err error){
	return
}

//TODO
func (p *Permiso) agregarPuntoDeAcceso(puntoDeAcceso funcionAPI) (err error){
	return
}

//TODO
func (p *Permiso) Registrar()(err error){
	return
}

func (p *Permiso) Revisar(u *Usuario)(tienePermiso bool, err error){
	//TODO
	return 
}

func (p *Permiso) Otorgar(u *Usuario)(err error){
	//TODO
	return
}