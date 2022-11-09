package main

import "fmt"

/*
	切片的零值为nil
	nil 切片的长度和容量为 0，并且没有底层数组。
*/
func main() {
	var s []int

	fmt.Println(s, len(s), cap(s))
	//todo //s := []int{}  时，s != nil  因为 := 为简洁赋值语句，分配内存空间，而 var s []int 只是声明，并没有分配内存，所以为nil
	if s == nil {
		fmt.Println("nil")
	}
}
