package persistencia

import (
	"bytes"
	"fmt"
	"strings"
)

//TODO
func validarNombre(nombre string) (err error) {
	return
}

//Convierte una *DefinicionDeBaseDeDatos en un query sql
func
parsearBaseDeDatos(b *DefinicionDeBaseDeDatos) (query string, err error) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(fmt.Sprintf("USE %s;\n", b.Nombre))
	for nombreTabla, definicionTabla := range b.Tablas {
		queryTabla := parsearTabla(definicionTabla)
		queryBuffer.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %s;\n", nombreTabla))
		queryBuffer.WriteString(fmt.Sprintf("CREATE TABLE %s (%s);\n", nombreTabla, queryTabla))
	}

	query = queryBuffer.String()

	return
}

//Convierte una *DefinicionDeTabla en un query sql correspondiente
func parsearTabla(t *DefinicionDeTabla) (query string) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("id INT UNIQUE NOT NULL AUTO_INCREMENT PRIMARY KEY")
	queryBuffer.WriteString(", ")
	for nombreColumna, definicionColumna := range t.Columnas {
		queryColumna := parsearColumna(definicionColumna)
		queryBuffer.WriteString(fmt.Sprintf("%s %s, ", nombreColumna, queryColumna))
	}
	query = queryBuffer.String()
	query = strings.TrimSuffix(query, ", ")
	return
}

//Convierte una *DefinicionDeColumna en un query sql correspondiente
func parsearColumna(c *definicionDeColumna) (query string) {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(c.TipoDeDatos.String())
	queryBuffer.WriteString(" ")
	if c.NotNull {
		queryBuffer.WriteString("NOT NULL")
		queryBuffer.WriteString(" ")
	}
	if c.Unique {
		queryBuffer.WriteString("UNIQUE")
		queryBuffer.WriteString(" ")
	}
	if c.DefaultValue != "" {
		queryBuffer.WriteString("DEFAULT")
		queryBuffer.WriteString(" ")
		queryBuffer.WriteString(c.DefaultValue)
	}
	query = queryBuffer.String()
	return
}
