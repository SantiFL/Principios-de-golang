package main

import (
	"fmt"      //imprime o devuelve texto
	"log"      //
	"net/http" // sirve para crear el servidor (peticiones http)

	// paquete que se usa para el enrutamiento
	"github.com/gorilla/mux" // usa la libreria gorilla desde la direccion de github sirve para crear peticiones para enrutamiento
)

func main() {
	router := mux.NewRouter().StrictSlash(true) // crea una nueva ruta, StrictSlash sirve para redireccionar la url
	router.HandleFunc("/", Index)               // crea la ruta index
	router.HandleFunc("/contacto", Contact)     // crea la ruta contacto

	server := http.ListenAndServe(":8080", router)  //sirve para hacer que el servidor escuche,y usa el puerto 8080 y el valor incial es el router
	log.Fatal(server)                               //es un metodo por si hay errores y los muestra por la consola
	fmt.Println("El servidor no se esta corriendo") // en caso de que no se genere la consulta al servidor crea un mensaje con una advertencia

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //llama ala funcion http,handlefunc sirve para crear una ruta,responseWrite enviar los mensajes al cliente HTTP,request peticiones del cliente http
	//fmt.Fprint(w, "hola desde un servidor local web con go") // escribe un mensaje en la pagina del servidor
}

// Index es parte de otro cosito
func Index(w http.ResponseWriter, r *http.Request) { // crea la funcion Index(pagina principal donde escribe y tira una peticion request)
	fmt.Fprintf(w, "este es la pagina principal") // imprime un mensaje en la pagina principal
}

// Contact es parte de otro coso
func Contact(w http.ResponseWriter, r *http.Request) { // crea una funcion Contact
	fmt.Fprintf(w, "este es la pagina de contacto")
}
