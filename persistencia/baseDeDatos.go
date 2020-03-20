package persistencia

import (
	"errors"
	"fmt"
)

type DefinicionDeBaseDeDatos struct {
	nombre string
	tablas map[string]*DefinicionDeTabla
}

type DefinicionDeTabla struct {
	columnas   map[string]*definicionDeColumna
	primaryKey string
}

//TODO: Foreign Key y órden de creación de tablas
type definicionDeColumna struct {
	tipoDeDatos  TipoDeDato
	notNull      bool
	unique       bool
	defaultValue string
}

type TipoDeDato string

func (t *TipoDeDato) String() string {
	return fmt.Sprintf("%s", *t)
}

const TINY_VARCHAR TipoDeDato = "VARCHAR(64)"
const VARCHAR TipoDeDato = "VARCHAR(128)"
const BIG_VARCHAR TipoDeDato = "VARCHAR(255)"
const TINY_INT TipoDeDato = "TINYINY"
const SMALL_INT TipoDeDato = "SMALLINT"
const INT TipoDeDato = "INT"
const BIG_INT TipoDeDato = "BIGINT"

func NuevaBaseDeDatos(nombre string) (baseDeDatos *DefinicionDeBaseDeDatos, err error) {
	//TODO: Revisar nombre inválido con regex
	baseDeDatos = new(DefinicionDeBaseDeDatos)
	if nombre == "" {
		err = errors.New("nombre de base de datos incorrecto")
	}
	if err == nil {
		baseDeDatos.nombre = nombre
		baseDeDatos.tablas = make(map[string]*DefinicionDeTabla)
	}
	return
}

//Agrega una definición de DefinicionDeTabla a la definición de base de datos
func (b *DefinicionDeBaseDeDatos) AgregarTabla(nombre string) (t *DefinicionDeTabla, err error) {
	err = b.validarTabla(nombre)

	if err == nil {
		b.tablas[nombre] = new(DefinicionDeTabla)
		t = b.tablas[nombre]
		t.columnas = make(map[string]*definicionDeColumna)
	}

	return
}

//valida que no exista una DefinicionDeTabla con el mismo nombre y que el nombre no sea ""
func (b *DefinicionDeBaseDeDatos) validarTabla(nombre string) (err error) {
	if nombre == "" {
		err = errors.New("nombre de DefinicionDeTabla incorrecto")
	}
	if err == nil {
		if _, existe := b.tablas[nombre]; existe {
			err = errors.New("ya está definida esa DefinicionDeTabla")
		}
	}
	return err
}

//Obtiene el puntero a una definición de tabla dentro de la definición de la base de datos
func (b *DefinicionDeBaseDeDatos) GetTabla(nombre string) (t *DefinicionDeTabla, err error) {
	var existe bool
	t, existe = b.tablas[nombre]
	if !existe {
		err = errors.New("error al buscar tabla en definición: no existe una tabla con ese nombre")
	}
	return
}

//Agrega una definicionDeColumna a una tabla
func (t *DefinicionDeTabla) AgregarColumna(nombre string, tipoDeDato TipoDeDato, defaultValue string, unique bool, notNull bool, primaryKey bool) (err error) {
	err = t.validarColumna(nombre, primaryKey)
	if err == nil {
		t.columnas[nombre] = &definicionDeColumna{
			tipoDeDatos:  tipoDeDato,
			defaultValue: defaultValue,
			unique:       unique,
			notNull:      notNull,
		}
		if primaryKey {
			t.primaryKey = nombre
		}
	}

	return
}

//Revisa que no exista una columna con el mismo nombre u otra columna como primaryKey
func (t *DefinicionDeTabla) validarColumna(nombre string, primaryKey bool) (err error) {
	var errorNombre string = ""
	var errorKey string = ""
	var separador string = ""
	if _, existe := t.columnas[nombre]; existe {
		errorNombre = "ya existe una columna con ese nombre"
	}
	if primaryKey && t.primaryKey != "" {
		errorKey = "ya hay otra primaryKey definida"
	}
	if errorKey != "" && errorNombre != "" {
		separador = " - "
	}
	if errorKey != "" || errorNombre != "" {
		err = errors.New(fmt.Sprintf("(%s%s%s)", errorNombre, separador, errorKey))
	}
	return
}
