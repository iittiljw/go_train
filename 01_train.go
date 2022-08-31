package main

import (
	"fmt"
)

const (
	asa = iota
	b   = 1 << (iota * 10)
	c   = 1 << (iota * 10)
	d   = 1 << (iota * 10)
	e   = 1 << (iota * 10)
	f   = 1 << (iota * 10)
	g   = 1 << (iota * 10)
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}
