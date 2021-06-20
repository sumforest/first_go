package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	fmt.Println()

	count := 1
	for count <= 10 {
		count += count
	}
	fmt.Println("count: ", count)

	// 无限循环
	// count3 := 1
	// for {
	// 	count3++
	// }
	// fmt.Println("count3:", count3)

	// for-eache range循环
	fmt.Println()
	strings := []string{"google", "runoob"}
	for index, val := range strings {
		fmt.Println("index: ", index, ", val: ", val)
	}

	fmt.Println()
	nums := [6]int{1, 2, 3}
	for index, val := range nums {
		fmt.Println("index: ", index, ", val: ", val)
	}
}
