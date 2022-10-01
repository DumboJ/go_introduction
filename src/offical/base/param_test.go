package main

import "testing"

func TestParam(t *testing.T) {
	t.Log(count(5, 14))
	t.Log(multi(10, 2))
}

//参数类型都指定
func count(x int, y int) int {
	return x + y
}

//连续多个参数类型相同，可以省略除最后一个参数外的类型
func multi(x, y int) int {
	return x * y
}
