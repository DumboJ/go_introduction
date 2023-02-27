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

//2.实现error方法
func (e *baseError) Error() string {
	return e.msg
}

//2.方法(两数相加)
func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a,b 不能为负数！")
	} else {
		return a + b, nil
	}
}

//4.改造add方法返回
func add_msg(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &baseError{
			1,
			"a,b 不能为负数！"}
	} else {
		return a + b, nil
	}
}
func main() {
	//普通未处理的error
	sum, ok := add(-2, 4)
	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(sum)
	}

	//带code编码的错误提示，使用断言
	sum1, msg := add_msg(-2, 3)
	if baseError, ok := msg.(*baseError); ok {
		fmt.Println("错误代码为：", baseError.code, "错误信息为：", baseError.msg)
	} else {
		fmt.Println(sum1)
	}
}
