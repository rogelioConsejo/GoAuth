package auth

import (
	"database/sql"
	"errors"
	"github.com/rogelioConsejo/Hecate/persistencia"
)

const tabla = "servicios"

type Servico struct {
	Nombre    string
	Direccion string
}

type ServiceEntity struct {
	Id_servicio uint
	Nombre      string
	Direccion   string
}

func (u ServiceEntity) GetId() uint {
	return u.Id_servicio
}

func RegistrarServicio(service Servico) (uint, error) {

	var err error
	var id uint

	if service.Nombre == "" && service.Direccion == "" {
		err = errors.New("error: nombre de usuario a registrar vacío")
	}

	if err == nil {
		entityRegistro := ServiceEntity{
			Nombre:    service.Nombre,
			Direccion: service.Direccion,
		}
		id, err = persistencia.RegistrarEnBaseDeDatos(entityRegistro, tabla)
	}
	return id, err
}

func BuscarServicio(nombre string) (s *ServiceEntity, id int, err error) {
	se := new(ServiceEntity)
	var row *sql.Row
	se.Nombre = nombre
	row, err = persistencia.BuscarUnoEnBaseDeDatos(*se, tabla)
	if err != nil {
		err = row.Scan(&s.Id_servicio, &s.Nombre, &s.Direccion)
	}
	return
}

func ActualizarServicio(id_servicio uint) (err error) {
	servicio := new(ServiceEntity)
	servicio.Id_servicio = id_servicio
	err = persistencia.ActualizarRegistroEnBaseDeDatos(servicio, tabla)
	return
}

func EliminarServicio(nombre string) (err error) {

	_, id, err := BuscarServicio(nombre)
	if err == nil {
		err = persistencia.BorrarEnBaseDeDatos(uint(id), TABLA_TOKENS)
	}
	return
}
