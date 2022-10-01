package main

/*go 的基本数据类型*/
import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("TYPE:% T ,value %v \n", ToBe, ToBe)
	fmt.Printf("TYPE:% T ,value %v \n", MaxInt, MaxInt)
	fmt.Printf("TYPE:% T ,value %v \n", z, z)
}
