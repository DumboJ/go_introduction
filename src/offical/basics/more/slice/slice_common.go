package main

import "fmt"

/*
* slice 动态数组更通用，可变数组
  slice 声明
		name [] T

	数组可以通过 [low:high]方式获取数组元素 slice分割数组是一个半开oper-half区间，左闭右开区间
*/
func main() {
	arrays := [6]int{1, 3, 4, 54, 23, 23}

	//slice 分割数组
	var a []int = arrays[1:3]
	fmt.Println(a)

	//slice 切片像是array数组的引用一样，修改切片元素的值，也会相应修改数组对应位置元素
	names := [4]string{
		"John",
		"Tom",
		"Anmy",
		"Linda",
	}
	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)

	x[1] = "QianJun"
	fmt.Println(x, y)
	fmt.Println(names)

	//分片定义
	numers := []int{12, 33, 13, 31, 33}
	bools := []bool{true, false, true, false, false}
	fmt.Println(numers, bools)

	//var stru = []struct {
	stru := []struct {
		i int
		b bool
	}{
		{2, true},
		{1, false},
		{12, true},
	}
	fmt.Println(stru)

	/*可变数组的分片时候的默认表示实例
	对于数组
		var a [10]int

	分片时下列操作一致
	[1:10]
	[:10]
	[0:]
	[:]
	*/
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(cap(s)) //7
	fmt.Println(s[:])
	fmt.Println(s[2:])
	s = s[1:4]
	fmt.Println(s) //[2 3 4]

	s = s[:2]
	fmt.Println(s) //[2 3]
	fmt.Println(len(s) /*2*/, cap(s) /*6*/)
	//todo 特殊slice use colon : 分片数组元素时，如果slice 表示 [0:],high 的缺省值是 len(s),即 s :=[]int{2,3} => len(s) = 2 => s[0:] == s[0:2]
	s = s[0:]
	fmt.Println(s) //[2 3]

	s = s[:1]
	fmt.Println(s)

}
