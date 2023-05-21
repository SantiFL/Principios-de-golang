package main

import (
	"fmt"
	"strings"
)

func main() {

	//a := "abc"
	b := "abcabcabc"

	fmt.Println(strings.Count(b, "abc"))
}
