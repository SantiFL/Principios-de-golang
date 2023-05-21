package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hola " + os.Args[1] + " por favor ingrese un numero")
	time.Sleep(time.Second * 3)

	primerdato, _ := strconv.Atoi(os.Args[2])
	segundodato, _ := strconv.Atoi(os.Args[3])

	suma := primerdato + segundodato

	fmt.Println("La suma es igual a :")
	fmt.Println(suma)
	sumar := 20
	resultado := 0
	if suma >= 20 {
		resultado = sumar + suma

		fmt.Println("la suma de los numero mas 20 es igual a:")
		fmt.Println(resultado)
	} else if suma < 20 {
		fmt.Println("como los datos son menor a 20 se restan ")
		resultado = primerdato - segundodato
		fmt.Println(resultado)
	}

	return
}
