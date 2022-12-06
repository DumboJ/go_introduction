package main

import (
	"fmt"
	"math"
)

type Location struct {
	x, y float64
}

/*
	Go 语言没有类，但可以通过type定义，并在方法参数中使用
*/
func (v Location) ABS() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	v := Location{3, 4}
	fmt.Println(v.ABS())
}
