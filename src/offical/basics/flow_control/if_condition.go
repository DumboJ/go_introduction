package main

import (
	"fmt"
	"math"
)

/*
	if 语句结构跟for类似，条件表达式的（）可省略 {}必须
*/
func main() {
	fmt.Println(sqrt(4), sqrt(-4))

	/*
		跟for语句一样，if 语句可以在条件表达式前执行一个简单语句  该语句声明的变量作用域，仅在if之内
	*/
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))

	/*
		if else demo : if中声明的短变量，可以在任何对应的else里使用
	*/
	fmt.Println(powWithElse(3, 3, 20))
}

/*计算数值开方后的值，如果是负数，先转正数后开方，并在结果后加”i”字符*/
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
	如果x的y次方小于n 返回 x的y次方计算值
	否则返回n的值
*/
func pow(x, y, n float64) float64 {
	if z := math.Pow(x, y); z < n {
		return z
	}
	return n
}

/*
	如果x的y次方小于n 返回 x的y次方计算值
	否则返回n的值
*/
func powWithElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g  >= %g \n", v, lim)
	}
	//这里之后不可以再使用if 内声明的变量 z
	return lim
}
