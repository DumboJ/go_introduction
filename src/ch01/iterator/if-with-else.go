package main

import (
	"fmt"
	"math"
)

/**
short variables in if statement scope
*/
func spow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(spow(2, 3, 10),
		spow(3, 3, 25),
	)
}
