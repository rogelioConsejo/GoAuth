package persistencia

import "testing"

func TestConfigurar(t *testing.T) {
	var direccion string = "localhost"
	var puerto int = 3038
	var usuario string = "hecate"
	var password string = "h3c4t3"
	c := &Conexion{
		DBdireccion: &direccion,
		DBpuerto:    &puerto,
		DBusuario:   &usuario,
		DBPassword:  &password,
	}
	err := Configurar(c)
	if err != nil{
		t.Errorf("error al configurar: %s\n", err.Error())
	}
}
