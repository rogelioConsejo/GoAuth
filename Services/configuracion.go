package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func ConfiguracionInicial() {
	serviceFile, err := ioutil.ReadFile("./Services/config.json")
	if err == nil {
		var datos map[string]interface{}
		json.Unmarshal(serviceFile, &datos)

		var valor string

		for key, value := range datos {
			valor += "{"
			for sk, sv := range value.(map[string]interface{}) {
				tipo := reflect.TypeOf(sv).String()
				if tipo == "map[string]interface {}" {
					valor += "\"" + sk + "\":{"
					for ssk, ssv := range sv.(map[string]interface{}) {
						valor += "\"" + ssk + "\":{"
						for sssk, sssv := range ssv.(map[string]interface{}) {
							valor += "\"" + sssk + "\":\"" + sssv.(string) + "\","
						}
						valor += "},"
					}
					valor += "}"
				} else {
					valor += "\"" + sk + "\":\"" + sv.(string) + "\","
				}
			}
			valor += "}"
			valor = strings.ReplaceAll(valor, ",}", "}")
			os.Setenv("go_"+key, valor)
			valor = ""
		}
	} else {
		log.Fatal("Error de lectura: ", err)
	}
}

func ExportarConfiguracion() {
	var llaves string
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.Contains(pair[0], "go_") {
			llaves += pair[0] + ","
		}
	}
	llaves += "-"
	llaves = strings.ReplaceAll(llaves, ",-", "")
	s := strings.Split(llaves, ",")
	var salida string
	salida = "{"
	for i := 0; i < len(s); i++ {
		k := strings.ReplaceAll(s[i], "go_", "")
		salida += "\"" + k + "\":"
		salida += os.Getenv(s[i]) + ","
	}
	salida += "-"
	salida = strings.ReplaceAll(salida, ",-", "")
	salida += "}"
	//fmt.Println(salida)
	err := ioutil.WriteFile("./Services/configExport.json", []byte(salida), 0644)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print( json.Marshal(salida),"\n")

}

func NuevaFuncion(nombreServicio string,nombreFuncion string,variables[]string)  {

}

func NuevoServicio(servicio string)  {
	serviceFile,_ := ioutil.ReadFile("./Services/config.json")
	serviceFile= []byte(strings.ReplaceAll(string(serviceFile), "}\n}", "},}"))
	runes := []rune(string(serviceFile))
	safeSubstring := string(runes[0:len(serviceFile)-3])+",\n"
	safeSubstring+="  "+servicio
	safeSubstring+="\n}"
	fmt.Println(safeSubstring)
	err := ioutil.WriteFile("./Services/config.json", []byte(safeSubstring), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
