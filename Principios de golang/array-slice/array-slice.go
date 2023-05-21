package main

import "fmt"

func main() {
	nombres := []string{
		"santiago",
		"Emanuel"}

	fmt.Println(nombres[0])
	fmt.Println(nombres[1])

	nombres = append(nombres, "sanson")

	fmt.Println(nombres[2])

	fmt.Println(len(nombres))

	fmt.Println(nombres[0:1])
}
