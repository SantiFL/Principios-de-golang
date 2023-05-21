package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	nuevotexto := os.Args[1]
	fichero, err := os.OpenFile("lenguajes.txt", os.O_APPEND, 0777)
	ShowError(err)

	escribir, err := fichero.WriteString(nuevotexto)
	fmt.Println(escribir)

	fichero.Close()

	texto, eroordefichero := ioutil.ReadFile("lenguajes.txt")
	ShowError(eroordefichero)

	fmt.Println(string(texto))

}

//ShowError sirve para saber si se encuentra un error
func ShowError(e error) {
	if e != nil {
		panic(e)
	}
}
