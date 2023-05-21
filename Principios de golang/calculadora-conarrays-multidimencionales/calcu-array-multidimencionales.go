package main

import "fmt"

func main() {
	var sumayresta [4][4]float32
	sumayresta[0][0] = 1
	sumayresta[0][1] = 2

	sumayresta[1][0] = 3
	sumayresta[1][1] = 4

	sumayresta[2][0] = 5
	sumayresta[2][1] = 6

	sumayresta[3][0] = 7
	sumayresta[3][1] = 8


	calculadorasuma(sumayresta[0][0], sumayresta[1][0], sumayresta[2][0], sumayresta[3][0])
	calculadoraresta(sumayresta[0][1], sumayresta[1][1], sumayresta[2][1], sumayresta[3][1])
	calculadoramultiplicacion(sumayresta[0][0], sumayresta[1][0], sumayresta[2][0], sumayresta[3][0])
	calculadoradivicion(sumayresta[0][1], sumayresta[1][1], sumayresta[2][1], sumayresta[3][1])
}
func calculos(n1 float32, n2 float32, n3 float32, n4 float32, op string) float32 {
	var res1 float32
	var res2 float32
	var resultadofinal float32

	if op == "+" {
		res1 = n1 + n2
		res2 = n3 + n4
		resultadofinal = res1 + res2

	}
	if op == "-" {
		res1 = n1 + n2 //6
		res2 = n3 + n4 //14
		resultadofinal = res1 - res2
	}
	if op == "*" {
		res1 = n1 * n2
		res2 = n3 * n4
		resultadofinal = res1 * res2
	}
	if op == "/" {
		res1 = n1 / n2
		res2 = n3 / n4
		resultadofinal = res1 / res2
	}

	return resultadofinal
}

func calculadorasuma(num1 float32, num2 float32, num3 float32, num4 float32) {
	fmt.Println("La suma es igual a: ")
	fmt.Println(calculos(num1, num2, num3, num4, "+"))
}

func calculadoraresta(numresta1 float32, numresta2 float32, numresta3 float32, numresta4 float32) {
	fmt.Println("La resta es igual a:")
	fmt.Println(calculos(numresta1, numresta2, numresta3, numresta4, "-"))
}

func calculadoramultiplicacion(nummul1 float32, nummul2 float32, nummul3 float32, nummul4 float32) {
	fmt.Println("La multiplicación es igual a:")
	fmt.Println(calculos(nummul1, nummul2, nummul3, nummul4, "*"))
}
func calculadoradivicion(numdiv1 float32, numdiv2 float32, numdiv3 float32, numdiv4 float32) {
	fmt.Println("La divicion es igual a: ")
	fmt.Println(calculos(numdiv1, numdiv2, numdiv3, numdiv4, "/"))
}
