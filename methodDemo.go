package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	var c Circle
	c.radius = 10.00
	fmt.Println("area = ", c.getArea())
}

// 这个方法属于Circle结构体中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
