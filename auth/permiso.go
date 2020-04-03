package auth

import (
	"errors"
)

type Permiso struct {
	nombre             string
	accionesPermitidas []*Accion
}

func CrearPermiso(nombre string) (p *Permiso, err error) {
	if nombre == "" {
		err = errors.New("no se puede crear un permiso sin nombre")
	}
	if err == nil {
		err = validarNombreDePermisoUnico(nombre)
	}
	p = new(Permiso)
	p.nombre = nombre
	p.accionesPermitidas = make([]*Accion, 0)

	return
}

//TODO
func validarNombreDePermisoUnico(nombre string) (err error) {
	err = errors.New("no implementado")
	return
}

//TODO
func RevisarPermiso(u *Usuario, p *Permiso)(tienePermiso bool, err error) {
	err = errors.New("no implementado")
	return
}

//TODO
func (p *Permiso) AgregarAccion(a *Accion) {
	return
}

//TODO
func (p *Permiso) Registrar() (err error) {
	return
}

func (p *Permiso) Otorgar(u *Usuario) (err error) {
	//TODO
	return
}
