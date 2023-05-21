package main

import "fmt"

type gorra struct {
	marca      string
	color      string
	disponible bool
	precio     float32
	plana      bool
}

type planta struct {
	nombre         string
	color          string
	precio         float32
	disponibilidad bool
}

func main() {
	var gorranegra = gorra{
		marca:      "nike",
		color:      "negra",
		precio:     25.20,
		plana:      false,
		disponible: true}
	fmt.Println(gorranegra)

	var plantavivero = planta{"Orquidia", "blanca", 2000.00, true}
	fmt.Println(plantavivero)

}
