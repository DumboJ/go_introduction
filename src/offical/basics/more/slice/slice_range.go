package main

import "fmt"

/*
	range 可以用于一个slice 或者map的for迭代，当range 用于slice上时，每次迭代对象会返回两个值：index和index对应的值
*/
var slice = []int{1, 22, 35, 55, 65, 125}

func main() {
	for i, v := range slice {
		fmt.Printf("index:%d--%d \n", i, v)
	}
}
