package main

import "fmt"

//todo 模仿error.NEW（text string）--实现errors接口

//1.errors类似接口 Success
type Success interface {
	Error() string
}

//2.定义errorString() 类似的类型
type successString struct {
	s string
}

//3.errorString结构体实现Success接口方法Error()
func (success *successString) Error() string {
	return success.s
}

//4.创建工厂方法 NEWSUCCESS()
func NEWSUCCESS(s string) Success {
	return &successString{s}
}
func main() {
	//5.调用
	a := NEWSUCCESS("factory method Success")
	fmt.Println(a)
}
