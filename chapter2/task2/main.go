package main

import (
	"fmt"
)

func main() {
	v := (1 == 1)
	c := (1 <= 2)
	d := (2 >= 1)
	x := (2 != 1)
	y := (1 < 2)
	o := (2 > 1)

	fmt.Println(v)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(o)
}
