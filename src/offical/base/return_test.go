package main

import "testing"

func Test_return(t *testing.T) {
	t.Log(returnOne("fire", "hole"))
	a, b := returnTwo("tok", "tik")
	t.Log(a, b)
}

func returnOne(x, y string) string {
	return x + "---" + y
}

//多值返回
func returnTwo(x, y string) (string, string) {
	return y, x

}
