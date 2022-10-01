package main

import (
	"fmt"
	"math"
)

/*
	类型转换
		表达式 T(v) 将值 v 转换为类型 T。
		var i int = 43
		var float32 = float32(i)
	Go 在不同类型的项之间赋值时需要显式转换
*/
func main() {
	var i, j int = 3, 4
	var f float64 = math.Sqrt(float64(i*i + j*j))
	var z = uint(f)
	fmt.Println(i, j, z)

}
