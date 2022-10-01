package main

import "fmt"

/*
	声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
		eg:
			var i int = 33
			var j = i
	当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度
		eg:
			i := 33 //int
			j := 0.234	//float64
			g := 0.125+0.5i //cmplx128

*/
func main() {
	i := 99
	g := 0.125 + 0.5i
	fmt.Printf(" this is type of %T ", i, g)
}
