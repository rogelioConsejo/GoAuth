package persistencia

import (
	"database/sql"
	"errors"
)

type BaseDeDatos struct {
	nombre string
	tablas map[string]*tabla
}

type tabla struct {
	columnas   map[string]*columna
	primaryKey string
}

//TODO: Foreign Key y órden de creación de tablas
type columna struct {
	tipoDeDatos  tipoDeDato
	notNull      bool
	unique       bool
	defaultValue string
}

type tipoDeDato string

func (t *tipoDeDato) String() string {
	return t.String()
}

const TINY_VARCHAR tipoDeDato = "VARCHAR(64)"
const VARCHAR tipoDeDato = "VARCHAR(128)"
const BIG_VARCHAR tipoDeDato = "VARCHAR(255)"
const TINY_INT tipoDeDato = "TINYINY"
const SMALL_INT tipoDeDato = "SMALLINT"
const INT tipoDeDato = "INT"
const BIG_INT tipoDeDato = "BIGINT"

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

//TODO
func (b *BaseDeDatos) RevisarErrores() (err error) {
	return
}

func (b *BaseDeDatos) Registrar(c *Conexion) (err error) {
	var baseDeDatos *sql.DB
	baseDeDatos, err = conectarABaseDeDatos(c)
	defer func() { err = cerrarConexion(baseDeDatos) }()

	//TODO

	return
}

//TODO: Corregir
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
