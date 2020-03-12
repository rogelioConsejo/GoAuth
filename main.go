package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/rogelioConsejo/Hecate/auth"
	"github.com/rogelioConsejo/Hecate/persistencia"
	"log"
	"net/http"
	"os"
)

const serverConfigPath = "serverConfig.conf"

type parametros struct {
	EsInstalacion     *bool                   `json:"EsInstalacion"`
	EsConfiguracion   *bool                   `json:"EsConfiguracion"`
	ConfiguracionDeBD persistencia.Conexion   `json:"ConfiguracionDeBD"`
	Configuracion     ConfiguracionDeServidor `json:"ConfiguracionDeServidor"`
}

type ConfiguracionDeServidor struct {
	DireccionDeServidor *string `json:"DireccionDeServidor"`
	PuertoDeServidor    *uint   `json:"PuertoDeServidor"`
}

func main() {
	log.Println("Programa servidor iniciado")
	var params parametros
	params = leerBanderas()

	//Comportamiento central del programa de despliegue de servidor API Gateway
	var err error
	if *params.EsInstalacion {
		err = persistencia.Instalar(params.ConfiguracionDeBD)
	} else if *params.EsConfiguracion {
		err = persistencia.CambiarConfiguracion(params.ConfiguracionDeBD)
	} else {
		err = correrServidor(*params.Configuracion.DireccionDeServidor, *params.Configuracion.PuertoDeServidor)
	}

	if err != nil {
		log.Printf("error de servidor: %s\n", err.Error())
	}
	defer log.Println("Cerrando servidor...")
}

//Definición de banderas
func leerBanderas() (p parametros) {
	p.EsInstalacion = flag.Bool("nuevo", false,
		"Indica que se quiere instalar por primera vez, borrando todas las tablas existentes")
	p.EsConfiguracion = flag.Bool("config", false,
		"Indica que se quiere configurar el programa")

	p.ConfiguracionDeBD.DBdireccion = flag.String("db", "", "La dirección de la base de datos")
	p.ConfiguracionDeBD.DBpuerto = flag.Int("dbport", 3306, "El puerto de la base de datos")
	p.ConfiguracionDeBD.DBusuario = flag.String("dbusr", "",
		"El nombre de usuario a usar para la Conexion a base de datos")
	p.ConfiguracionDeBD.DBPassword = flag.String("dbpass", "",
		"El password a usar para la Conexion a base de datos")

	p.Configuracion.DireccionDeServidor = flag.String("d", "localhost",
		"La direccion en donde será accesible el servidor, se debe definir también un puerto")
	p.Configuracion.PuertoDeServidor = flag.Uint("p", 8080,
		"El puerto desde donde será accesible el servidor, se debe definir también una dirección")

	flag.Parse()

	return
}

func correrServidor(direccion string, puerto uint) (err error) {
	var config ConfiguracionDeServidor
	config, err = getConfiguracionDeServidor(direccion, puerto)

	if err == nil {
		log.Print("Ejecutando servidor... ")
		http.HandleFunc("/", handler)
		if *config.PuertoDeServidor >= 0 {
			puerto = 8080
		}
		addr := *config.DireccionDeServidor + ":" + fmt.Sprint(*config.PuertoDeServidor)
		log.Printf("Ubicación: %s\n", addr)
		err = http.ListenAndServe(addr, nil)
	}

	return
}

func getConfiguracionDeServidor(direccion string, puerto uint) (config ConfiguracionDeServidor, err error) {
	var file *os.File
	var configurar bool = direccion != "" && puerto > 0

	_, err = os.Stat(serverConfigPath)
	if configurar {
		log.Printf("Configurando servidor: %s:%d\n", direccion, puerto)
		config, err = configurarServidor(direccion, puerto)
	} else if os.IsNotExist(err) {
		err = errors.New("el servidor no está configurado, usar -h para más ayuda")
	} else {
		log.Println("Leyendo configuración del servidor")
		if err == nil {
			file, err = os.Open(serverConfigPath)
			defer func() { err = file.Close() }()
		}
		if err == nil {
			var configJSON []byte
			_, err = file.Read(configJSON)
			err = json.Unmarshal(configJSON, config)
		}
	}

	return
}

func configurarServidor(direccion string, puerto uint) (config ConfiguracionDeServidor, err error) {
	if puerto == 0 || puerto > 65535 {
		err = errors.New("puerto inválido")
	}

	var file *os.File
	var configuracionJSON []byte
	if err == nil {
		config.DireccionDeServidor = &direccion
		config.PuertoDeServidor = &puerto
		file, err = os.OpenFile(serverConfigPath, os.O_RDWR|os.O_CREATE, 0755)
	}
	if err == nil {
		configuracionJSON, err = json.Marshal(config)
	}
	if err == nil {
		_, err = file.Write(configuracionJSON)
	}

	if err != nil {
		nuevoErr := fmt.Sprintf("error al configurar servidor: %s\n", err)
		err = errors.New(nuevoErr)
	}
	return
}

//PROGRAMA PRINCIPAL DEL SERVIDOR
func handler(response http.ResponseWriter, request *http.Request) {
	var usuario *auth.Usuario
	var tienePermiso bool
	var accionARealizar accion
	var err error
	var email string
	var pass string
	usuario, err = auth.RevisarCredenciales(email, pass)
	if err == nil {
		accionARealizar, err = parsearPeticion(request)
	}

	if err == nil {
		tienePermiso, err = accionARealizar.permiso.Revisar(usuario)
	}

	if tienePermiso && err == nil {
		log.Printf("Petición (%s): %s\n", usuario.GetEmail(), accionARealizar.getIdentificador())
		var resultado *resultado
		resultado, err = accionARealizar.do(usuario)
		if resultado != nil {
			log.Printf("Resultado (%s): %s -> %s\n",
				usuario.GetEmail(), accionARealizar.getIdentificador(), resultado.getMensaje())
		}
	} else if !tienePermiso && err == nil {
		log.Printf("ALERTA: usuario %s intentó realizar una acción sin permiso: %s\n", usuario.GetEmail(),
			accionARealizar.getIdentificador())
	}

	if err != nil{
		log.Printf("error en API Gateway: %s\n", err.Error())
	}
}

//TODO: Implementar
func parsearPeticion(request *http.Request) (accion accion, err error) {
	return
}
