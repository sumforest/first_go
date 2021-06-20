package main

import "fmt"

var a, b int = 1, 2

// 通常用于声名全局变量
var (
	c int
	d int
)

var e, f int = 3, 4
var g, h = 12, "hello"

// 通常用于函数体中声名变量
//  i, j := 'a','b'

func main() {

	fmt.Println(a, b, c, d, e, f, g, h)
}
