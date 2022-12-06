package main

import (
	"fmt"
	"math"
)

/*
todo 在这个例子中，我们看到了一个MyFloat带有Abs方法的数字类型。
	您只能声明一个带有接收者的方法，该接收者的类型定义在与该方法相同的包中。您不能使用类型在另一个包（包括内置类型，例如int）中定义的接收器来声明方法
*/
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	f := MyFloat(math.Sqrt2)
	fmt.Println(f.Abs())
}
