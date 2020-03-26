package main

import (
	"fmt"
	"github.com/rogelioConsejo/Hecate/auth"
)

type Permiso struct {
	id uint32
	nombre string
	puntosDeAcceso funciones
}

type funciones []funcionAPI

type funcionAPI struct {
	servicio   service
	verbo      *HTTPverb
	recurso    string
}

type HTTPverb string

const (
	GET  HTTPverb = "GET"
	POST HTTPverb = "POST"
	PUT  HTTPverb = "PUT"
	DELETE HTTPverb = "DELETE"
)

func (verb *HTTPverb)String()string{
	return fmt.Sprintf("%s", *verb)
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

func (p *Permiso) Revisar(u *auth.Usuario)(tienePermiso bool, err error){
	//TODO
	return 
}

func (p *Permiso) Otorgar(u *auth.Usuario)(err error){
	//TODO
	return
}