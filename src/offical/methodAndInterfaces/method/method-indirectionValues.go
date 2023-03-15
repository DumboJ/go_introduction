package main

//summary : 方法定义，不限制方法的接收者为值或指针，可以编译阶段转换。但是函数定义，传递参数值和指针需要区分
import (
	"fmt"
	"math"
)

type Coordinate struct {
	x, y float64
}

//参数为值类型的函数
func abs(c Coordinate) float64 {
	return math.Sqrt(c.x*c.x + c.y*c.y)
}

// 值为接收者的方法
func (c Coordinate) abdFun() float64 {
	return math.Sqrt(c.x*c.x + c.y*c.y)
}
func main() {
	//结论：值为接收者的方法，调用时，接收者可以是值 也可以是 指针
	//		但是方法参数为值时，参数必须为指定类型的值，不可以是指针
	c := Coordinate{3, 4}
	fmt.Println(c.abdFun())
	fmt.Println(abs(c))

	c1 := &Coordinate{3, 4}
	fmt.Println(c1.abdFun())
	fmt.Println(abs(c1))  //compiler error : Cannot use 'c1' (type *Coordinate) as the type Coordinate
	fmt.Println(abs(*c1)) //compiler success
}
