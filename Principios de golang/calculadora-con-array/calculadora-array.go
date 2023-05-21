package main

import "fmt"

func main() {
	var suma = "+"
	var resta = "-"
	var multiplicacion = "*"
	var divicion = "/"
	var numeros [4]float32
	numeros[0] = 1
	numeros[1] = 2
	numeros[2] = 3
	numeros[3] = 4

	fmt.Println(numeros)
	fmt.Println("la suma es igual:")
	fmt.Println(operacion(numeros[0], numeros[1], numeros[2], numeros[3], suma))
	fmt.Println("la resta es igual:")
	fmt.Println(operacion(numeros[0], numeros[1], numeros[2], numeros[3], resta))
	fmt.Println("la multiplicacion es igual:")
	fmt.Println(operacion(numeros[0], numeros[1], numeros[2], numeros[3], multiplicacion))
	fmt.Println("la divicion es igual:")
	fmt.Println(operacion(numeros[0], numeros[1], numeros[2], numeros[3], divicion))
}

func operacion(n1 float32, n2 float32, n3 float32, n4 float32, operador string) float32 {
	var res1 float32

	if operador == "+" {
		res1 = n1 + n2 + n3 + n4

		return res1
	}
	if operador == "-" {
		res1 = n1 - n2 - n3 - n4

		return res1
	}

	if operador == "*" {
		res1 = n1 * n2 * n3 * n4

		return res1
	}

	if operador == "/" {
		res1 = n1 / n2 / n3 / n4

		return res1
	}
	return res1

}
