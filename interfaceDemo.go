package main

import "fmt"

func main() {
	iphone := new(Iphone)
	iphone.call()

	android := new(Android)
	android.call()
}

type Phone interface {
	// 无入参，无返回值
	call()
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("Iphone call you")
}

type Android struct {
}

func (android Android) call() {
	fmt.Println("Android call you")
}
