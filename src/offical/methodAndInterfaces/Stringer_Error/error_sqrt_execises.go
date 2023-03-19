package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//parse e to float , or it will spark Callback hell
	return fmt.Sprintf("cannot Sqrt negative number:%g", float64(e))
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
func main() {
	fmt.Println(sqrt(-2))
	fmt.Println(sqrt(2))
}
