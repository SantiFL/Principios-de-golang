package main

import (
	"fmt"
	"time"
)

func main() {
	momento := time.Now()
	hoy := momento.Weekday()
	switch hoy {
	/*
		case 0:
			fmt.Println("hoy es lunes")
		case 1:
			fmt.Println("hoy es martes")
		case 2:
			fmt.Println("hoy es miercoles")
		case 3:
			fmt.Println("hoy es juevez")
	*/
	case 0:
		fmt.Println("hoy es viernes")
	case 1:
		fmt.Println("hoy es sabado")
	case 2:
		fmt.Println("hoy es domingo")

	}
}
