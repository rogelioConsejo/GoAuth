package persistencia

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestInstalar(t *testing.T) {
	usuario := "root"
	password := ""
	direccion := "localhost"
	nombre := "hecate"
	c := &ConfiguracionDeConexion{
		DBusuario:   &usuario,
		DBPassword:  &password,
		DBdireccion: &direccion,
		DBnombre:    &nombre,
	}
	db, err := NuevaBaseDeDatos("hecate")
	tabla, err := db.AgregarTabla("usuarios")
	if err != nil {
		log.Print("1")
		t.Error(err.Error())
	}
	err = tabla.AgregarColumna("mail", VARCHAR, "", true, true, false)
	if err != nil {
		log.Print("2")
		t.Error(err.Error())
	}
	err = tabla.AgregarColumna("passwordHash", VARCHAR, "", false, true, false)
	if err != nil {
		log.Print("3")
		t.Error(err.Error())
	}
	jsonDB, err := json.Marshal(db)
	log.Println(fmt.Sprintf("%s\n", jsonDB))
	err = Instalar(c,db)
	if err != nil {
		log.Print("4")
		t.Error(err.Error())
	}

}
