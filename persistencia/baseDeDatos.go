package persistencia

import (
	"database/sql"
	"errors"
)

type BaseDeDatos struct {
	nombre string
	tablas map[string]tabla
}

type tabla struct {
	nombre   string
	columnas columnas
}

type columnas map[string]string

func NuevaBaseDeDatos(nombre string) (baseDeDatos *BaseDeDatos, err error) {
	//TODO: Revisar nombre inválido con regex
	if nombre == "" {
		err = errors.New("nombre de base de datos incorrecto")
	}
	if err == nil {
		baseDeDatos.nombre = nombre
	}
	return
}

func (b *BaseDeDatos) Registrar(c *Conexion) (err error){
	var baseDeDatos *sql.DB
	baseDeDatos, err = conectarABaseDeDatos(c)
	defer func() {err = cerrarConexion(baseDeDatos)}()

	//TODO

	return
}

func (b *BaseDeDatos) AgregarTabla(nombre string, columnas map[string]string) (err error) {
	err = b.validarTabla(nombre, columnas)

	if err == nil {
		var t tabla = tabla{nombre: nombre, columnas: columnas}
		b.tablas[nombre] = t
	}

	return
}

func (b *BaseDeDatos) validarTabla(nombre string, columnas map[string]string) (err error) {
	//TODO: Revisar nombre inválido con regex
	if nombre == "" {
		err = errors.New("nombre de tabla incorrecto")
	}
	if len(columnas) <= 0 {
		err = errors.New("no se definieron columnas al agregar tabla")
	}
	if err == nil {
		if _, existe := b.tablas[nombre]; existe {
			err = errors.New("ya está definida esa tabla")
		}
	}
	return err
}