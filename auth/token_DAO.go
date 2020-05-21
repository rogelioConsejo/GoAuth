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

func guardarToken(t *Token) (err error) {
	entity := t.Entity()
	_, err = persistencia.RegistrarEnBaseDeDatos(*entity, TABLA_TOKENS)

	return
}

func borrarToken(id uint) (err error) {
	err = persistencia.BorrarEnBaseDeDatos(id, TABLA_TOKENS)
	return
}

func buscarToken(codigo string) (t *Token, id uint, err error) {
	e := new(tokenEntity)
	t = new(Token)
	var row *sql.Row
	e.Codigo = codigo
	row, err = persistencia.BuscarUnoEnBaseDeDatos(*e,TABLA_TOKENS)
	if err == nil {
		err = row.Scan(&e.Id, &e.Codigo, &e.Usuario)
	}
	if err == nil {
		id = e.Id
		t.Codigo = e.Codigo
		t.usr, err = leerUsuarioEnBaseDeDatos(e.Usuario)
	}

	return
}
