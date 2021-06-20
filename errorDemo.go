package main

import "fmt"

func main() {
	if result, msg := devide(100, 10); msg == "" {
		fmt.Println("result = ", result)
	}

	if _, msg := devide(100, 0); msg != "" {
		fmt.Println("errormsg: ", msg)
	}
}

type DivideError struct {
	// 除数
	dividee int
	// 被除数
	divider int
}

// 实现error接口
func (de *DivideError) Error() string {
	strFomat := `
	不能够执行出发，除数为0
	被除数：%d
	除数：0
	`
	return fmt.Sprintf(strFomat, de.dividee)
}

// 除法函数
func devide(dividee int, divider int) (result int, errormsg string) {
	if divider == 0 {
		data := DivideError{dividee, divider}
		errormsg = data.Error()
		return
	} else {
		return dividee / divider, ""
	}
}
