package main

import "net/http"

func main()  {
	//TODO: banderas

	//TODO: Set-up base de datos SQL con banderas - usando un archivo de configuración

	//TODO: Correr servidor con dirección configurable por banderas


	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil{
		//TODO: Loggear error de montado de servidor
	}
}

func handler(response http.ResponseWriter, request *http.Request){

}