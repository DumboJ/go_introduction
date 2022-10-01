package main

import "fmt"

/*
	基础数据类型的初始值
		即零值是：
			数值类型为 0，
			布尔类型为 false，
			字符串为 ""（空字符串）
*/
func main() {
	var i int
	var j float32
	var k bool
	var r string
	fmt.Printf("%v %v %v %v \n", i, j, k, r)
}
