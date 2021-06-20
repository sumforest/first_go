package main

import "fmt"

func main() {

	// 创建通道
	c := make(chan int)
	s := []int{1, 2, 3, 4, 5, 6}

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println("x = ", x, "y = ", y)

	fmt.Println("-------------- buffer channel -----------")
	// 创建带缓冲的通道
	ch := make(chan int, 100)
	ch <- 2
	ch <- 3
	fmt.Println(<-ch, ",", <-ch)
}

// 计算数组之和并且发送到通道
func sum(s []int, c chan int) {
	var sum int
	for _, val := range s {
		sum += val
	}

	// 发送到通道
	c <- sum
}
