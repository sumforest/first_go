package main

import "fmt"

func main() {
	// 切片
	var sliceDemo = make([]int, 3, 5)
	// 数组
	var arr = []int{1, 2, 3}
	// 切片实质就是动态数组
	fmt.Printf("sliceDemo的类型为:%T\n", sliceDemo)
	fmt.Printf("sliceDemo的类型为:%T\n", arr)
	printSlice(sliceDemo)

	fmt.Println()
	// nil空切片
	var emptySlice []int
	printSlice(emptySlice)
	if nil == emptySlice {
		fmt.Println("切片emptySlice为空")
	}
}

func printSlice(x []int) {
	fmt.Printf("len = %d,capcity = %d,slice = %v\n", len(x), cap(x), x)
}
