package main

import "fmt"

func main() {
	// 声名一个变量并且赋初值
	var wesite = "Runoob"
	fmt.Println(wesite)

	// 生命变量使用默认值
	var a int
	fmt.Println(a)

	var isInit bool
	fmt.Println(isInit)

	fmt.Println()
	// 相当于 var website2 string
	// website2 = "www.sumforest.cn"
	website2 := "www.sumforest.cn"
	fmt.Println(website2)
}
