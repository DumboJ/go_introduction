package ch02_test

import "testing"

const (
	Monday = iota + 1
	Thuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
const (
	Readable   = 1 << iota //001
	Writable               //010
	Executable             //100
)

func TestConst(t *testing.T) {
	t.Log(Monday, Thuesday, Wednesday, Thursday)
}
func TestBitCalculate(t *testing.T) {
	//and := 7 //0111
	and := 1                        //1
	t.Log(and&Readable == Readable, //001 & 001  true
		and&Writable == Writable,     //001 & 010  false
		and&Executable == Executable) //001 & 100 false

}
