package main

import "fmt"

func main() {
	getNumber := getSquence()

	fmt.Println(getNumber())
	fmt.Println(getNumber())
	fmt.Println(getNumber())

	fmt.Println("----------")
	// 创建一个新的函数对象
	getNumber2 := getSquence()
	fmt.Println(getNumber2())
	fmt.Println(getNumber2())
}

func getSquence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
