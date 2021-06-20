package main

import "fmt"

func main() {
	c := make(chan int, 10)

	go feibonacci(cap(c), c)

	for val := range c {
		fmt.Print(val, ",")
	}
}

// 返回第n项斐波那次数列
func feibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 此处不调用close方法将在main中循环range中一直阻塞造成死锁
	close(c)
}
