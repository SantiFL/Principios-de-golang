package main

import "fmt"

func main() {
	pantalon("rojo", "largo", "nike", "otros") // se declara en la funcion pantalon que posea los datos siguientes
}

func pantalon(caracteristicas ...string) { //el ... sirve para uno o muchos parametros de una forma abreviada en este caso son muchos parametros de tipo string
	for _, caracteristica := range caracteristicas { //el _ es para no pasarle ninguna variable, range es para recorrer por todas las caracteristicas
		fmt.Println(caracteristica) // imprime las caracteristicas
	}
}
