package main

import (
	"fmt"
	"math"
)
import "math/cmplx"
import "time"

func main() {
	fmt.Println(math.Abs(-2))
	fmt.Println(cmplx.Sqrt(-3 + 21i))
	fmt.Println("now time is ", time.Now())
}
