package main

import "testing"

//var语句用于声明一个变量列表 跟函数的参数列表一样，类型在最后
var python, Java bool

func TestVal(t *testing.T) {
	//var 语句可以出现在包或函数级别
	var i int
	t.Log(i, python, Java)
}
