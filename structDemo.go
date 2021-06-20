package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	id      int
}

func main() {
	fmt.Println(Book{"go 语言", "Google 巨佬", "计算机科学与技术", 1})
	fmt.Println(Book{title: "java", author: "Sun", subject: "计算机科学与与技术", id: 2})
	fmt.Println(Book{title: "python", author: "忘记了哪个巨佬", id: 3})
}
