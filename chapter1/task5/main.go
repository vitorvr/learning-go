package main

import (
	"fmt"
)

type mytype int
var x mytype
var y int

func main() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v\t%T\n", y, y)
}