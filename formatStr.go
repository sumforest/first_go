package main

import (
	"fmt"
)

func main(){
	var code = 123
	var date = "2021年6月17日"
	var url = "Code=%d&endDate=%s"
	var result = fmt.Sprintf(url,code,date)
	fmt.Println(result)
}