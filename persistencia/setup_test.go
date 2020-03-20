package persistencia

import (
	"log"
	"testing"
)

func TestInstalar(t *testing.T) {
	c := &ConfiguracionDeConexion{
		DBdireccion: "localhost",
		DBpuerto:    3306,
		DBusuario:   "hecate",
		DBPassword:  "h3c4t3",
		DBnombre:    "hecate",
	}
	db, err := NuevaBaseDeDatos("hecate")
	tabla, err := db.AgregarTabla("usuarios")
	if err != nil {
		log.Print("1")
		t.Error(err.Error())
	}
	err = tabla.AgregarColumna("mail", VARCHAR, "", true, true, true)
	if err != nil {
		log.Print("2")
		t.Error(err.Error())
	}
	err = tabla.AgregarColumna("pass", VARCHAR, "", false, true, false)
	if err != nil {
		log.Print("3")
		t.Error(err.Error())
	}
	err = Instalar(c,db)
	if err != nil {
		log.Print("4")
		t.Error(err.Error())
	}

}
