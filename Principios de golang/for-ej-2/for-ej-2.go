package main

import "fmt"

func main() {
	for i := 0; i <= 15; i++ {
		if i%2 == 0 {
			fmt.Println("El numero es par")
			fmt.Println(i)
		} else if i%2 != 0 {
			fmt.Println("El numero es impar")
			fmt.Println(i)
		}
	}
}
