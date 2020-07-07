package persistencia

import (
	"fmt"
	"testing"
)

func TestConfigurar(t *testing.T) {
	usuario := "su"
	password := "Halo4ygokugt"
	direccion := "localhost"
	nombre := "hecate"
	c := &ConfiguracionDeConexion{
		DBusuario:   &usuario,
		DBPassword:  &password,
		DBdireccion: &direccion,
		DBnombre:    &nombre,
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

	*c.DBnombre = ""
	err = Configurar(c)
	if err == nil{
		t.Error("Nombre de BD vacío no detectado")
	}
	*c.DBnombre = "hecate"
	*c.DBusuario = ""
	err = Configurar(c)
	if err == nil{
		t.Error("Nombre de usuario vacío no detectado")
	}


}

func TestConfigurar2(t *testing.T) {
	config, err := getConfiguracion()
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("%+v\n", config)
	fmt.Printf("%+v\n", new(ConfiguracionDeConexion))
}