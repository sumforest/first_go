package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32

	var ptr *int

	// %T打印变量的类型
	fmt.Printf("a的类型为：%T\n", a)
	fmt.Printf("b的类型为：%T\n", b)
	fmt.Printf("c的类型为：%T\n", c)

	ptr = &a
	fmt.Printf("a的值为：%d\n", a)
	fmt.Printf("*ptr的值为：%d", *ptr)

}
