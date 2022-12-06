package main

import (
	"fmt"
	"math"
)

type Vertexa struct {
	x, y float64
}

//返回开方的计算结果
func (v Vertexa) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

//指针类型入参，并在该方法中通过指针修改Vertexa中元素的值
//如果入参不使用 *T 指针形式，则修改后的值不影响最终结果
//由于方法通常需要修改其接收者，因此指针接收者比值接收者更常见。
//对于值接收器，该方法对原始值Scale的副本进行操作
func (v Vertexa) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}
func main() {
	//声明 Vertexa 赋值 x=3 y=4
	v := Vertexa{3, 4}
	//通过读取指针，并修改值
	v.Scale(10)
	fmt.Println(v.Abs())

}
