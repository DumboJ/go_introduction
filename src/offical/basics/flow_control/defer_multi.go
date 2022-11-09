package main

import "fmt"

/*
	defer推迟的函数会被压入栈内存中，外层函数调用完后，被推迟的栈中函数会按LIFO（后进先出）的顺序调用
*/
func main() {
	fmt.Print("counting ")
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, "-")
	}
	fmt.Println("done")
}
