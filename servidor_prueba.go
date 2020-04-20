package main

import (
	"fmt"
	"net/http"
)

func manejador(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola, %s, ¡Este es el servidor!", r.URL.Path)
}

func main(){
	http.HandleFunc("/", manejador)
	fmt.Println("El servidor se encuentra en ejecución")
	http.ListenAndServe(":8080", nil)
}