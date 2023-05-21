package main

import "fmt"

func main() {
	fmt.Print("pedidos 1 ---->")
	fmt.Println(gorras(45, "$"))
	fmt.Println("--------------------------")
	fmt.Print("pedidos 2")
	fmt.Println(gorras(10, "U$S"))
}

func gorras(pedidos float32, moneda string) (string, float32, string) {

	precio := func() float32 {
		return pedidos * 7

	}

	return "el valor de las gorras pedidas son: ", precio(), moneda

}
