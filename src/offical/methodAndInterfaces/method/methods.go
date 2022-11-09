package main

import (
	"fmt"
	"math"
)

type Location struct {
	x, y float64
}

func (v Location) ABS() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	v := Location{3, 4}
	fmt.Println(v.ABS())
}
