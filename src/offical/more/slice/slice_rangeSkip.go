package main

import "fmt"

/*
	针对 range slice 迭代时，可以通过 _ 跳过 index 或 value
	当只需要index 时，也可以直接省略第二个变量
*/
func main() {
	slice := make([]int, 10)
	for i := range slice {
		//todo 移位index个无符号值
		slice[i] = i << uint(i)
	}
	for _, v := range slice {
		fmt.Printf("%d \n", v)
	}
}
