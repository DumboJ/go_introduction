package main

import "testing"

func testv(t *testing.T) {
	/*在函数中，简洁赋值语句 := 可以用在类型明确的地方代替 var声明
	函数外的每个语句都必须以关键字（var,func）开始，因此  := 不能在函数外使用
	*/
	var h int = 3
	r := 9
	z := "vari"
	t.Log(h, r, z)
}
