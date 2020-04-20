package auth

import (
	"database/sql"
	"proyectos_sys/Hecate/persistencia"
)

const(
	TB_TOKENS = "tokens"
)

type tbTokens struct {
	Id uint
	Codigo string
	Usuario uint
}

func (e tbTokens) GetId() uint {
	return e.Usuario
}

func buscaUsuarioToken(codigo string) (t *token, usuario uint, err error) {
	e:= new(tbTokens)
	t = new(token)
	var row *sql.Row
	e.Codigo = codigo
	row, err = persistencia.BuscarUnoEnBaseDeDatos(*e, TB_TOKENS)
	if err == nil {
		err = row.Scan(e.Usuario)
	}
	if err == nil{
		t.usr, err = leerUsuarioEnBaseDeDatos(e.Usuario)
	}
	return
}