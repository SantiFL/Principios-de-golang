package main

import (
	"fmt"      // libreria de formatos(textos)
	"log"      // libreria de loggin
	"net/http" // libreria para crear un seridor web con las fuciones http

	"github.com/gorilla/mux" //libreria para crear peticiones de enrutamiento
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/Aprendido", Aprendido)
	router.HandleFunc("/Terminado", Terminado)
	router.HandleFunc("/En curso", 	Encurso)

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
	fmt.Printf("Ocurrio un problema en el server")
}

//Index es la pagina principal
func Index(I http.ResponseWriter, n *http.Request) {
	fmt.Fprintf(I, "Pagina Principal")
}

//Aprendiendo es la pagina donde muestro lo que aprendí hasta ahora
func Aprendido(A http.ResponseWriter, p *http.Request) {
	fmt.Fprintf(A, "Pagina donde muestro lo que se ")
}

//Terminado es la pagina de cursos o carreras terminadas
func Terminado(T http.ResponseWriter, e *http.Request) {
	fmt.Fprintf(T, "Estos son los cursos termiandos")
}
