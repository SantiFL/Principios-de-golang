package main

import "fmt"

func main() {
	var peliculas [3][2]string
	peliculas[0][0] = "rapidos y furiosos"
	peliculas[0][1] = "rapidos y furiosos 2"

	peliculas[1][0] = "rapidos y furiosos 3 reto tokio"
	peliculas[1][1] = "rapidos y furiosos 4"

	peliculas[2][0] = "rapidos y furiosos 5"
	peliculas[2][1] = "rapidos y furiosos 6"

	fmt.Println(peliculas[0][0])
}
