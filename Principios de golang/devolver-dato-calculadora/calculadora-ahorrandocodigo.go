package main

import "fmt"

func main() {
	var numero int = 10
	var numero2 int = 10

	fmt.Println(delvolvertexto(""))
	fmt.Println(operacion(numero, numero2, "+"))
}

func operacion(num1 int, num2 int, op string) int {
	var resultado1 int
	if op == "+" {
		resultado1 = num1 + num2
	}
	return resultado1
}

func delvolvertexto(dato1 string) string {
	dato1 = "La suma es igual:"
	return dato1
}
