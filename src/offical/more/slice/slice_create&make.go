package main

import (
	"fmt"
)

/*
	可以使用内置make函数创建切片；创建动态大小的数组的方式。
	a := make([]int, 5) // len(a)=5
	上述语句表示：使用make函数分配一个归零的数组并返回一个引用该数组的切片
	make创建动态数组时，第三个参数指定容量
*/
func main() {
	a := make([]int, 5)
	printMakeSlice("a", a)
	//make([]int ,数组长度,数组容量)
	b := make([]int, 0, 6)
	printMakeSlice("b", b)
	//对len = 0的数组赋值，会报异常 index out of range [0] with length 0
	//b[0] = 1

	//todo 对于len(b)=0,但有实际cap值的对象，可以重新切分slice指定slice[low,high]
	c := b[:2]
	/* todo 大于引用数组 cap()+1的slice 将会报异常 slice bounds out of range [:7] with capacity 6
	c1 := b[:7]
	printMakeSlice("c1", c1)*/
	printMakeSlice("c", c)

	d := b[2:5]
	printMakeSlice("d", d)

	f := b[2:2]
	printMakeSlice("f", f)
}
func printMakeSlice(num string, s []int) {
	fmt.Printf("%s cap:%d len:%d  %v \n", num, cap(s), len(s), s)
}
