//3 arrays de los cuales 2 vas a ser multidimencionales     (completado)
//agregar 2 palabras o numero a uno de los arrays           (completado)
//que muestre datos (nombre,apellido,dni,fecha de nacimiento)(completado)
//sumar dos numeros de los propios arrays                    (completado)
//mostrar un struc                                            (completado)
//por ultimo un mensaje de bienvenida con un temporizador     (completado)
//mostrar cantidad de datos de unos de los array
package main

import (
	"fmt"
	"time"
)

func main() {
	datospersonales := []string{
		"juan",
		"casas"}
	datospersonales = append(datospersonales, "alfredo")
	datospersonales = append(datospersonales, "rodriguez")
	time.Sleep(time.Second * 10)
	fmt.Println("hola bienbenido al programa propio echo por mi Santi")
	fmt.Println(len(datospersonales))
	fmt.Println("Nombre:")
	fmt.Println(datospersonales[0])
	fmt.Println("apellido")
	fmt.Println(datospersonales[1])

	var datospersonales2 [1][2]int
	datospersonales2[0][0] = 15780697
	datospersonales2[0][1] = 25122020
	fmt.Println("dni")
	fmt.Println(datospersonales2[0][0])
	fmt.Println("fecha de nacimiento")
	fmt.Println(datospersonales2[0][1])

	var numeros [2][2]float32
	numeros[0][0] = 10
	numeros[0][1] = 10

	fmt.Println("la suma es igual a: ")
	fmt.Println(numeros[0][0] + numeros[0][1])

	type gorra struct {
		marca      string
		color      string
		disponible bool
		precio     float32
		plana      bool
	}

	var gorranegra = gorra{
		marca:      "nike",
		color:      "negra",
		precio:     25.20,
		plana:      false,
		disponible: true}
	fmt.Println(gorranegra)
}
