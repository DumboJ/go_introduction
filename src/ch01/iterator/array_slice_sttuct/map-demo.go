package main

import (
	"fmt"
)

//init
func main() {
	//1.make
	userMap := make(map[string]uint8)

	//2.字面量方式(也可以空声明)
	GOODS := map[string]float64{"iphone": 6300, "pencil": 3.2}

	Computer := map[string]string{}

	//增
	userMap["DumboJ"] = 25

	GOODS["MacBook pro"] = 7800

	Computer["联想"] = "拯救者R7000P"
	Computer["华硕"] = "冒险家"

	//删
	delete(userMap, "Lucky")
	//改
	GOODS["MacBook pro"] = 7650
	//查
	age, ok := userMap["DumboJ"]
	if ok {
		fmt.Println(age)
	}

	//遍历 for key,value := range map{}
	for k, v := range Computer {
		fmt.Println("笔记本品牌", k, "。笔记本型号：", v)
	}

	fmt.Println("Computer 容量：", len(Computer))
}
