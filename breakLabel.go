package main

import "fmt"

func main() {
	fmt.Println("----------- break -----------")
	for i := 1; i < 3; i++ {
		fmt.Println("i = ", i)
		for j := 10; j < 13; j++ {
			fmt.Println("j = ", j)
			break
		}
	}

	fmt.Println("----------- break-label -----------")
flag:
	for i := 1; i < 3; i++ {
		fmt.Println("i = ", i)
		for j := 10; j < 13; j++ {
			fmt.Println("j = ", j)
			break flag
		}
	}
}
