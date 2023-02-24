package main

import "fmt"

/*
	1.GO中函数都会从属于一个包
	2.函数名大写，不同包可以调用
	3.函数名小写，只能在相同包中调用
*/
func main() {
	fmt.Print("a")
	//todo println 等特殊 从属builtin包。该包提供了go预先声明的函数、变量等的文档
	println()
}
