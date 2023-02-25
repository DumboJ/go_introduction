package main

import (
	"errors"
	"fmt"
)

//1.声明结构体
type baseError struct {
	code int8
	msg  string
}

//2.方法(两数相加)
func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a,b 不能为负数！")
	} else {
		return a + b, nil
	}
}
func main() {
	sum, ok := add(-2, 4)
	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(sum)
	}
}
