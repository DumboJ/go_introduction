package main

import "testing"

func Testvarinit(t *testing.T) {
	//变量声明可以包含初始值，每个变量对应一个
	var i, j int = 2, 3
	t.Log(i, j)
	var c, python, java = true, false, "no"
	t.Log(c, python, java)
}
