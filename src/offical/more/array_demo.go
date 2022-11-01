package main

import "fmt"

func main() {
	//declare array  --   name ['size']type
	//数组的长度是其类型的一部分，所以数组不能被调整大小，但是不用担心，Go提供了一种处理数组的便捷方法
	var sayHello [2]string
	sayHello[0] = "Hello"
	sayHello[1] = "world"
	fmt.Println(sayHello[0], sayHello[1])
	fmt.Println(sayHello)

	prime := [5]int{1, 2, 3, 4, 5}
	fmt.Println(prime)
}
