package main

import "fmt"

func main() {

	a := 3000

	var ptr *int
	var pptr **int

	ptr = &a
	pptr = &ptr

	fmt.Println("a的值为：", a)
	fmt.Println("ptr的值为：", *ptr)
	fmt.Println("指针的指针变量的值为：", **pptr)
}
