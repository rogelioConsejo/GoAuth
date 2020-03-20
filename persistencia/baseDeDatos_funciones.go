package persistencia

import (
	"bytes"
	"fmt"
)

//TODO
func validarNombre(nombre string) (err error) {
	return
}

//Convierte una *DefinicionDeBaseDeDatos en un query sql
func parsearBaseDeDatos(b *DefinicionDeBaseDeDatos) (query string, err error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(fmt.Sprintf("CREATE DATABASE %s;\n", b.nombre))
	for nombreTabla, definicionTabla := range b.tablas {
		queryTabla := parsearTabla(definicionTabla)
		queryBuffer.WriteString(fmt.Sprintf("CREATE TABLE %s (%s);\n", nombreTabla, queryTabla))
	}

	query = queryBuffer.String()

	return
}

//Convierte una *DefinicionDeTabla en un query sql correspondiente
func parsearTabla(t *DefinicionDeTabla) (query string) {
	var queryBuffer bytes.Buffer
	for nombreColumna, definicionColumna := range t.columnas {
		queryColumna := parsearColumna(definicionColumna)
		queryBuffer.WriteString(fmt.Sprintf("%s %s, ", nombreColumna, queryColumna))
	}

	return
}

//Convierte una *DefinicionDeColumna en un query sql correspondiente
func parsearColumna(c *definicionDeColumna) (query string) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(c.tipoDeDatos.String())
	queryBuffer.WriteString(" ")
	if c.notNull {
		queryBuffer.WriteString("NOT NULL")
		queryBuffer.WriteString(" ")
	}
	if c.unique {
		queryBuffer.WriteString("UNIQUE")
		queryBuffer.WriteString(" ")
	}
	if c.defaultValue != "" {
		queryBuffer.WriteString("DEFAULT")
		queryBuffer.WriteString(" ")
		queryBuffer.WriteString(c.defaultValue)
	}
	query = queryBuffer.String()
	return
}
