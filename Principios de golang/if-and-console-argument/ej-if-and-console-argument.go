package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("hi and welcome to my program")
	fmt.Println("hi " + os.Args[1] + " welcome")

	yearold, _ := strconv.Atoi(os.Args[2])
	year, _ := strconv.Atoi(os.Args[3])

	if yearold >= 18 {
		fmt.Println("you are of age")
	}

	fmt.Println("the year is: ")
	fmt.Println(year)
}
