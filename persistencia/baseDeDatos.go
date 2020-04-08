package persistencia

import (
	"errors"
	"fmt"
)

type DefinicionDeBaseDeDatos struct {
	Nombre string
	Tablas map[string]*DefinicionDeTabla
}

//TODO: Implementar PrimaryKey diferente de id
type DefinicionDeTabla struct {
	Columnas   map[string]*definicionDeColumna
	//PrimaryKey string
}

//TODO: Foreign Key y órden de creación de Tablas
type definicionDeColumna struct {
	TipoDeDatos  TipoDeDato
	NotNull      bool
	Unique       bool
	DefaultValue string
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
	//TODO: Revisar Nombre inválido con regex
	baseDeDatos = new(DefinicionDeBaseDeDatos)
	if nombre == "" {
		err = errors.New("Nombre de base de datos incorrecto")
	}
	if err == nil {
		baseDeDatos.Nombre = nombre
		baseDeDatos.Tablas = make(map[string]*DefinicionDeTabla)
	}
	return
}

//Agrega una definición de DefinicionDeTabla a la definición de base de datos
func (b *DefinicionDeBaseDeDatos) AgregarTabla(nombre string) (t *DefinicionDeTabla, err error) {
	err = b.validarTabla(nombre)

	if err == nil {
		b.Tablas[nombre] = new(DefinicionDeTabla)
		t = b.Tablas[nombre]
		t.Columnas = make(map[string]*definicionDeColumna)
	}

	return
}

//valida que no exista una DefinicionDeTabla con el mismo Nombre y que el Nombre no sea ""
func (b *DefinicionDeBaseDeDatos) validarTabla(nombre string) (err error) {
	if nombre == "" {
		err = errors.New("Nombre de DefinicionDeTabla incorrecto")
	}
	if err == nil {
		if _, existe := b.Tablas[nombre]; existe {
			err = errors.New("ya está definida esa DefinicionDeTabla")
		}
	}
	return err
}

//Obtiene el puntero a una definición de tabla dentro de la definición de la base de datos
func (b *DefinicionDeBaseDeDatos) GetTabla(nombre string) (t *DefinicionDeTabla, err error) {
	var existe bool
	t, existe = b.Tablas[nombre]
	if !existe {
		err = errors.New("error al buscar tabla en definición: no existe una tabla con ese Nombre")
	}
	return
}

//Agrega una definicionDeColumna a una tabla
func (t *DefinicionDeTabla) AgregarColumna(nombre string, tipoDeDato TipoDeDato, defaultValue string, unique bool, notNull bool, primaryKey bool) (err error) {
	err = t.validarColumna(nombre, primaryKey)
	if err == nil {
		t.Columnas[nombre] = &definicionDeColumna{
			TipoDeDatos:  tipoDeDato,
			DefaultValue: defaultValue,
			Unique:       unique,
			NotNull:      notNull,
		}
	}

	return
}

//Revisa que no exista una columna con el mismo Nombre u otra columna como PrimaryKey
func (t *DefinicionDeTabla) validarColumna(nombre string, primaryKey bool) (err error) {
	var errorNombre string = ""
	var errorKey string = ""
	var separador string = ""
	if _, existe := t.Columnas[nombre]; existe {
		errorNombre = "ya existe una columna con ese Nombre"
	}
	if errorKey != "" && errorNombre != "" {
		separador = " - "
	}
	if errorKey != "" || errorNombre != "" {
		err = errors.New(fmt.Sprintf("(%s%s%s)", errorNombre, separador, errorKey))
	}
	return
}
