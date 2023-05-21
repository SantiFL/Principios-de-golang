package main

import "fmt"

func main() {

	var numero1 float32 = 10
	var numero2 float32 = 5
	fmt.Println("Calculadora 1")
	calculadora1(numero1, numero2)

	fmt.Println("---------------------------------")

	var numero3 float32 = 5
	var numero4 float32 = 10

	fmt.Println("calculadora 2")
	calculadora2(numero3, numero4)

}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32
	if op == "+" {
		resultado = n1 + n2
	}
	if op == "-" {
		resultado = n1 - n2
	}
	if op == "*" {
		resultado = n1 * n2
	}
	if op == "/" {
		resultado = n1 / n2
	}
	return resultado
}

func calculadora1(numero1 float32, numero2 float32) {
	fmt.Print("La suma es igual: ")
	fmt.Println(operacion(numero1, numero2, "+"))
	fmt.Print("La resta es igual: ")
	fmt.Println(operacion(numero1, numero2, "-"))
	fmt.Print("La mutiplicacion es igual: ")
	fmt.Println(operacion(numero1, numero2, "*"))
	fmt.Print("La divicion es igual: ")
	fmt.Println(operacion(numero1, numero2, "/"))

}

func calculadora2(numero3 float32, numero4 float32) {
	fmt.Print("La suma es igual: ")
	fmt.Println(operacion(numero3, numero4, "+"))
	fmt.Print("La resta es igual: ")
	fmt.Println(operacion(numero3, numero4, "-"))
	fmt.Print("La mutiplicacion es igual: ")
	fmt.Println(operacion(numero3, numero4, "*"))
	fmt.Print("La divicion es igual: ")
	fmt.Println(operacion(numero3, numero4, "/"))

}
