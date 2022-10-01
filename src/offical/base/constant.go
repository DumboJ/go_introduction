package main

import "fmt"

/*
	常量声明：
		常量声明与变量声明类似，关键字更改为 constant
			变量声明 ：eg  var i int = 3
			常量声明 ：eg constant pi float32 = 3.144

		常量可以是字符串、字符、布尔、数值
	!!!!!!!常量不能用短变量方式声明 即 :=
*/
const PI = 3.14

func main() {
	const valid = true
	fmt.Println("GO rules?", valid)

	const world = "world"
	fmt.Println("Hello ", world)

	fmt.Println("PI value is ", PI, "nice math")
}
