package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rogelioConsejo/Hecate/persistencia"
	"log"
	"net/http"
	"os"
)

const serverConfigPath = "serverConfig.conf"

type parametros struct {
	EsInstalacion     *bool                                 `json:"EsInstalacion"`
	EsConfiguracion   *bool                                 `json:"EsConfiguracion"`
	ConfiguracionDeBD *persistencia.ConfiguracionDeConexion `json:"ConfiguracionDeBD"`
	Configuracion     *ConfiguracionDeServidor               `json:"ConfiguracionDeServidor"`
}

type ConfiguracionDeServidor struct {
	DireccionDeServidor *string `json:"DireccionDeServidor"`
	PuertoDeServidor    *uint   `json:"PuertoDeServidor"`
}

//TODO: Que no se guarde la configuración del servidor nueva si ya hay una guardada y no hay bandera (default)
func main() {

	var params parametros
	params = leerBanderas()

	//Comportamiento central del programa de despliegue de servidor API Gateway
	var err error
	if *params.EsInstalacion {
		var definicionBD *persistencia.DefinicionDeBaseDeDatos
		definicionBD, err = generarDefinicionDeBaseDeDatos()
		if err == nil {
			err = persistencia.Instalar(params.ConfiguracionDeBD, definicionBD)
		}
	} else if *params.EsConfiguracion {
		err = persistencia.Configurar(params.ConfiguracionDeBD)
	} else {
		err = correrServidor(*params.Configuracion.DireccionDeServidor, *params.Configuracion.PuertoDeServidor)
	}

	if err != nil {
		log.Printf("error de ejecución: %s\n", err.Error())
	}
	defer log.Println("Cerrando servidor...")
}

func correrServidor(direccion string, puerto uint) (err error) {
	var config ConfiguracionDeServidor
	config, err = getConfiguracionDeServidor(direccion, puerto)

	if err == nil {
		log.Print("Ejecutando servidor... ")
		http.HandleFunc("/", handler)
		http.HandleFunc("/usr/", usrHandler)
		http.HandleFunc("/login/", loginHandler)
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
			r := json.NewDecoder(file)
			err = r.Decode(&config)
		}
	}

	if err != nil {
		nuevoError := fmt.Sprintf("error al obtener configuración de servidor: %s\n", err.Error())
		err = errors.New(nuevoError)
	}

	return
}

func configurarServidor(direccion string, puerto uint) (config ConfiguracionDeServidor, err error) {
	if puerto == 0 || puerto > 65535 {
		err = errors.New("puerto inválido")
	}

	var file *os.File

	if err == nil {
		config.DireccionDeServidor = &direccion
		config.PuertoDeServidor = &puerto
		file, err = os.OpenFile(serverConfigPath, os.O_RDWR|os.O_CREATE, 0755)
	}

	if err == nil {
		c := json.NewEncoder(file)
		err = c.Encode(config)
	}

	if err != nil {
		nuevoErr := fmt.Sprintf("error al configurar servidor: %s\n", err)
		err = errors.New(nuevoErr)
	}
	return
}
