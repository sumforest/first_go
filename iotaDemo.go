package main

import "fmt"

const (
	a = 1 << iota
	b = 3 << iota
	c
	d
)

func main() {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
}
