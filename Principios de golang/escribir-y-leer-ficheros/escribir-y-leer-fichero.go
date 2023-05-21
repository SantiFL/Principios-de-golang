package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	nuevotexto := os.Args[1] + "\n"

	fichero, err2 := os.OpenFile("lenguajes.txt", os.O_APPEND, 077)
	MostrarError(err2)
	escribir, err3 := fichero.WriteString(nuevotexto)
	MostrarError(err3)
	fmt.Println(escribir)
	fichero.Close()

	leer, err := ioutil.ReadFile("lenguajes.txt")
	MostrarError(err)
	fmt.Println(string(leer))
}

func MostrarError(e error) {
	if e != nil {
		panic(e)
	}

}
