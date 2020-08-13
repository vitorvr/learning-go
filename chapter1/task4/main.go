package main

import (
	"fmt"
)

type mytype int
var x mytype

func main() {
	fmt.Printf("%v\t%T\n", x, x)
	x = 42
	fmt.Println(x)
}