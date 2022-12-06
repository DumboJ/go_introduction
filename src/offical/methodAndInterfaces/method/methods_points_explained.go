package main

import (
	"fmt"
	"math"
)

type Param struct {
	x, y float64
}

//修改指针值
func Scale(v *Param, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

//开方
func Abs(p Param) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func main() {
	p := Param{3, 4}
	Scale(&p, 10)
	fmt.Println(Abs(p))
}
