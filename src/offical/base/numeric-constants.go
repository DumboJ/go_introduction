package main

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	BIG = 1 << 100

	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	SMALL = BIG >> 99
)

func main() {
	/*数值常量是高精度的值
	一个未指定类型的数值常量由上下文来决定类型
	*/
	fmt.Println(needInt(SMALL))
	fmt.Println(needFloat(SMALL))
	fmt.Println(needFloat(BIG))
}
func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
