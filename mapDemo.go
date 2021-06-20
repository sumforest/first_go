package main

import "fmt"

func main() {
	var countryMap map[string]string
	// 初始化map
	countryMap = make(map[string]string)

	countryMap["China"] = "北京"
	countryMap["Janpan"] = "东京"
	countryMap["Korean"] = "首尔"
	countryMap["Russia"] = "莫斯科"

	// 输出map
	for country := range countryMap {
		fmt.Println(country, "，首都是：", countryMap[country])
	}

	// 查看元素是否存在集合中
	capital, ok := countryMap["America"]
	if ok {
		fmt.Println("美国存在map中，首都是：", capital)
	} else {
		fmt.Println("美国不存在map中")
	}

	fmt.Println("---------- 删除map ---------")
	delete(countryMap, "Janpan")
	// 输出删除后的map
	for country := range countryMap {
		fmt.Println(country, "，首都是：", countryMap[country])
	}

}
