package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var sum int
	// 用不上index，使用_丢掉
	for _, val := range arr {
		sum += val
	}
	fmt.Println("arr每项之和：", sum)

	kvs := map[string]string{"a": "aa", "b": "bb", "c": "cc"}
	for k, v := range kvs {
		fmt.Println("k: ", k, "v: ", v)
	}
}
