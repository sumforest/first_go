package main

import (
	"fmt"
	"unsafe"
)

// 函数必须在const里面是使用
const (
	e = "abc"
	f = len(e)
	g = unsafe.Sizeof(e)
)

func main() {
	const LEIGH, WIDTH = 20, 10

	var area int

	const a, b, c = 1, false, "hello"

	area = LEIGH * WIDTH
	fmt.Println(area)
	fmt.Println()

	fmt.Println(a, b, c)

	println(e, f, g)
}
