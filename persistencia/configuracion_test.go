package persistencia

import (
	"fmt"
	"testing"
)

func TestConfigurar(t *testing.T) {
	c := &ConfiguracionDeConexion{
		DBusuario:   "hecate",
		DBPassword:  "h3c4t3",
		DBdireccion: "localhost",
		DBnombre:    "hecate",
	}
	err := Configurar(c)
	if err != nil {
		t.Errorf("error al configurar: %s\n", err.Error())
	}

	c2, err := getConfiguracion()
	if err != nil {
		t.Errorf("error al obtener configuración: %s\n", err.Error())
	} else if !(c2.DBusuario==c.DBusuario&&c2.DBdireccion==c.DBdireccion&&c2.DBPassword==c.DBPassword&&c2.DBnombre==c.DBnombre) {
		t.Error("configuración obtenida no corresponde con la registrada")
	}

	c.DBnombre = ""
	err = Configurar(c)
	if err == nil{
		t.Error("nombre de BD vacío no detectado")
	}
	c.DBnombre = "hecate"
	c.DBusuario = ""
	err = Configurar(c)
	if err == nil{
		t.Error("nombre de usuario vacío no detectado")
	}


}

func TestConfigurar2(t *testing.T) {
	config, err := getConfiguracion()
	println(err.Error())
	fmt.Printf("%+v", config)
	fmt.Printf("%+v", new(ConfiguracionDeConexion))
}