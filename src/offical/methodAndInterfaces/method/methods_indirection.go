package main

import "fmt"

type Vertx struct {
	X, Y float64
}

//带有指针接收者的方法在调用时将值或者指针作为接收者
func (v *Vertx) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//带有指针参数的函数必须采用指针传参调用
func ScaleFunc(v *Vertx, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	v := Vertx{3, 4}
	//当v是一个值，而不是指针时，调用Scale方法时，相当于GO会自动将语句解释为(&v).Scale(2)的形式，因为该方法具有指针接收器
	v.Scale(2)
	//必须指针传参
	ScaleFunc(&v, 10)

	//ScaleFunc(v,10)		--compile:error

	p := &Vertx{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)
	fmt.Println(v)
	fmt.Println(p)
}
