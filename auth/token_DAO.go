package auth

import (
	"database/sql"
	"github.com/rogelioConsejo/Hecate/persistencia"
)

const (
	TABLA_TOKENS = "tokens"
)

type tokenEntity struct {
	Id      uint
	Codigo  string
	Usuario uint
}

func (e tokenEntity) GetId() uint {
	return e.Id
}

func guardarToken(t *token) (err error) {
	entity := t.Entity()
	_, err = persistencia.RegistrarEnBaseDeDatos(*entity, TABLA_TOKENS)

	return
}

func borrarToken(id uint) (err error) {
	err = persistencia.BorrarEnBaseDeDatos(id, TABLA_TOKENS)
	return
}

func buscarToken(codigo string) (t *token, id uint, err error) {
	e := new(tokenEntity)
	t = new(token)
	var row *sql.Row
	e.Codigo = codigo
	row, err = persistencia.BuscarUnoEnBaseDeDatos(*e,TABLA_TOKENS)
	if err == nil {
		err = row.Scan(&e.Id, &e.Codigo, &e.Usuario)
	}
	if err == nil {
		id = e.Id
		t.codigo = e.Codigo
		t.usr, err = leerUsuarioEnBaseDeDatos(e.Usuario)
	}

	return
}
