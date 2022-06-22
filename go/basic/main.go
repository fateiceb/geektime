package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
	d
)
const (
	e = iota
	f
	g
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(e, f, g)
}
