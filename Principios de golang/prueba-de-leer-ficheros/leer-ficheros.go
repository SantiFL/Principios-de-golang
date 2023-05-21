package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("lector:")

	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")
	ShowError(errorDeFichero)

	fmt.Println(string(texto))
}

// ShowError es una prubea por si el lector de error muestra un error
func ShowError(e error) {
	if e != nil {
		panic(e)
	}
}
