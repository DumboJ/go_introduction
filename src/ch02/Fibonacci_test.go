package ch02_test

import "testing"

func TestFebonacci(t *testing.T) {
	//var a = 1
	//var b = 17
	a, b := 0, 1
	t.Log(a)
	t.Log("\r\n")
	for i := 0; i < 6; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = a + tmp
	}
}
