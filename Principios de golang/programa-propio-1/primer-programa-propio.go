package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	const oquis = "blanco"
	const monitor = "negro"
	const fechdenac = 31121997
	fechdenacstring := strconv.Itoa(fechdenac)
	var impresora string = "negra"
	var teclado string = "verde"
	var lapicera string = "negra"
	var fechdehoy int = 2892020
	fechdehoystring := strconv.Itoa(fechdehoy)
	off := "naranja"
	programaconsola := "cygwin"
	profesor := " Victor Robles"
	time.Sleep(time.Second * 5)
	fmt.Println("El color de la planta es: " + oquis)
	time.Sleep(time.Second * 7)
	fmt.Println("El color del monitor es: " + monitor)
	time.Sleep(time.Second * 9)
	fmt.Println("Fecha de nacimiento:" + fechdenacstring)
	time.Sleep(time.Second * 11)
	fmt.Println("la impresora es color:" + impresora)
	time.Sleep(time.Second * 11)
	fmt.Println("El teclado es:" + teclado)
	time.Sleep(time.Second * 9)
	fmt.Println("La laciera es de color: " + lapicera)
	time.Sleep(time.Second * 7)
	fmt.Println("La fecha de hoy es:" + fechdehoystring)
	time.Sleep(time.Second * 5)
	fmt.Println("El off es color: " + off)
	time.Sleep(time.Second * 3)
	fmt.Println("El programa para mostrar en consola es: " + programaconsola)
	time.Sleep(time.Second * 1)
	fmt.Println("El profesor de este curso es: " + profesor)

}
