package auth

import (
	"database/sql"
	"proyectos_sys/Hecate/persistencia"
)

const (
	TABLA_ACCIONES = "acciones"
)

type accionEntity struct {
	Id uint
	nombreDelServicio string
	verboHTTP HttpVerb
	identificadorDeRecurso uri
	body requestBody
}

func (a accionEntity) GetId() uint {
	return a.Id
}

func guardarAccion(a *Accion)(err error){
	entity := a.Entity()
	_, err = persistencia.RegistrarEnBaseDeDatos(*entity, TABLA_ACCIONES)

	return
}

func borrarAccion(id uint)(err error){
	err = persistencia.BorrarEnBaseDeDatos(id, TABLA_ACCIONES)
	return
}

func buscarAccion(nombreServicio string)(a *Accion, id uint, err error){
	e := new(accionEntity)
	a = new(Accion)
	var row *sql.Row
	a.nombreDelServicio = nombreServicio
	row, err = persistencia.BuscarUnoEnBaseDeDatos(*e, TABLA_ACCIONES)
	if err == nil{
		err = row.Scan(&e.Id, &e.nombreDelServicio, &e.verboHTTP, &e.identificadorDeRecurso, &e.body)
	}
	if err == nil{
		id = e.Id
		a.nombreDelServicio = e.nombreDelServicio
		a.verboHTTP = e.verboHTTP
		a.identificadorDeRecurso = e.identificadorDeRecurso
		a.body = e.body
	}
	return
}