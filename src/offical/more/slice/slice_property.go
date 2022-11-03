package main

import "fmt"

func main() {
	s := []int{1, 3, 4, 5, 6, 9, 11}
	//原始数组 len:7 cap=7 [1 3 4 5 6 9 11]
	printSlice(s)
	//切分 slice 长度为0  len:0 cap=7 []
	printSlice(s[:0])
	//切分数组范围内  len:4 cap=7 [1 3 4 5]
	printSlice(s[:4])
	//切分移除数组前两位  len:5 cap=5 [4 5 6 9 11]
	printSlice(s[2:])
}
func printSlice(s []int) {
	fmt.Printf("len:%d cap=%d %v\n", len(s), cap(s), s)
}
