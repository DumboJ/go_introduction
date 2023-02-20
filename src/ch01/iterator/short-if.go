package main

import (
	"fmt"
	"math"
)

/**
short variables in if statement scope
*/
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(pow(2, 3, 10),
		pow(3, 3, 25),
	)
}
