package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("Los lenguajes son: ")

	fichero, MostrarError := ioutil.ReadFile("lenguajes.txt")
	ShowError(MostrarError)

	fmt.Println(string(fichero))
}

// ShowError crea una funcion para ver si hay algun error en el fichero
func ShowError(e error) {
	if e != nil {
		panic(e)
	}
}
