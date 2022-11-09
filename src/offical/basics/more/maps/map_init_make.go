package main

import "fmt"

type Location struct {
	Lat, Long float32
}

//map 声明 name map[key type] value type
var m map[string]Location

func main() {
	//声明后的map 为nil ，nil的map没有键值，也不可以添加键值
	if m == nil {
		fmt.Println(m)
	}
	//make 方法初始化已声明的map类型备用
	m = make(map[string]Location)
	m["Bob Lily"] = Location{
		8.22,
		3.022,
	}
	fmt.Println(m["Bob Lily"])
	var n = make(map[string]float32)
	n["gg"] = 33.23
	fmt.Println(n["gg"])
}
