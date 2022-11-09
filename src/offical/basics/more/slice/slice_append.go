package main

import "fmt"

func main() {
	/*
		将新元素附加到切片是很​​常见的，因此 Go 提供了一个内置 append函数。 内置包的文档append描述了.

		func append(s []T, vs ...T) []T
		的第一个参数s是appendtype 的切片，T其余的是 T要附加到切片的值。

		的结果值append是一个切片，其中包含原始切片的所有元素以及提供的值。

		如果 的后备数组s太小而无法容纳所有给定值，则将分配更大的数组。返回的切片将指向新分配的数组。
	*/
	var s []int
	printSliceInfo(s)

	//append适用于nil
	s = append(s, 0)
	printSliceInfo(s)

	//切片可以按需增长
	s = append(s, 1)
	printSliceInfo(s)

	//可以同时增加多个元素
	s = append(s, 1, 3, 4, 5)
	printSliceInfo(s)

}

func printSliceInfo(s []int) {
	fmt.Printf("len:%d, cap:%d, %v \n", len(s), cap(s), s)
}
