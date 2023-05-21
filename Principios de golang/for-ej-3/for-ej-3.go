package main

import (
	"fmt"
)

func main() {

	fmt.Println("Estas son las series que vi")
	fmt.Println("ingresa un numero para saber que serie estoy viendo ahora ")
	seriesquevi := []string{"Hause", "The big bang theory", "naruto", "Demons layer", "Dragon ball", "the god of high school"}
	for p := 0; p < len(seriesquevi); p++ {

		fmt.Println(seriesquevi[p], p)
	}

}
