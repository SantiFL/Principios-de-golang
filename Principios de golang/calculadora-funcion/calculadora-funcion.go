package main

import "fmt"

func main() {
	var numero1 float32 = 10
	var numero2 float32 = 20
	var numero3 float32 = 30
	var numeropararestar float32 = 60
	calculadora1(numero1, numero2, numero3, numeropararestar)
}

func operacion(numero1 float32, numero2 float32, numero3 float32, numresta float32, op string) float32 {
	var resultado1 float32
	var resultado2 float32
	if op == "+" {
		resultado1 = numero1 + numero2
		resultado2 = resultado1 + numero3

	}

	if op == "-" {
		resultado1 = numero1 + numero2 + numero3
		resultado2 = resultado1 - numresta
	}

	if op == "*" {
		resultado1 = numero1 * numero2 * numero3 * numresta
		resultado2 = resultado1
	}

	if op == "/" {
		resultado1 = numero1 / numero2 / numero3 / numresta
		resultado2 = resultado1
	}

	return resultado2
}

func calculadora1(n1 float32, n2 float32, n3 float32, n4 float32) {
	fmt.Println("La suma es igual a: ")
	fmt.Println(operacion(n1, n2, n3, n4, "+"))
	fmt.Println("La resta es igual a: ")
	fmt.Println(operacion(n1, n2, n3, n4, "-"))
	fmt.Println("La multiplicacion es igual a: ")
	fmt.Println(operacion(n1, n2, n3, n4, "*"))
	fmt.Println("La divicion es igual a: ")
	fmt.Println(operacion(n1, n2, n3, n4, "/"))
}
