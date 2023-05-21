package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Ej 3 in go")

	numero, _ := strconv.Atoi(os.Args[1])
	numero2, _ := strconv.Atoi(os.Args[2])
	signo := (os.Args[3])

	if numero%2 == 0 {
		fmt.Println("el primero numero ingresado es par")
	} else {
		fmt.Println("el primero numero ingresado es impar")
	}
	if numero2%2 != 0 {
		fmt.Println("El segundo numero ingresado es impar")
	} else {
		fmt.Println("El segundo numero ingresado es par ")
	}

	if signo == "+" {
		resultado := numero + numero2
		fmt.Println("El resultado de el primero numero + el segundo es :")
		fmt.Println(resultado)
	} else if signo == "-" {
		resulatado2 := numero - numero2
		fmt.Println("El resultado de la resta esta igual a :")
		fmt.Println(resulatado2)
	}

}
