package main

import "fmt"

func main() {
	var peliculas [3]string
	peliculas[0] = "forrest gomp"
	peliculas[1] = "batman"
	peliculas[2] = "iron man"

	fmt.Println(peliculas)
	fmt.Println(peliculas[0])
	fmt.Println(peliculas[1])
	fmt.Println(peliculas[2])

	peliculas2 := [3]string{
		"iron man",
		"iron man2",
		"iron man3"}

	fmt.Println(peliculas2)

}
