package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//字符串不可变的字节序列 所以可以直接转换成字节切片
	say := "Hello世界"
	bytes := []byte(say)
	fmt.Println(say)
	fmt.Println(bytes[0])
	//一个汉字占3个字节，所以，hello + 世界 = 5 + 3*2 = 11
	fmt.Println(len(say))

	//想将一个汉字按照一个长度计算，可以使用函数
	fmt.Println(utf8.RuneCountInString(say))

	//for range 在循环处理字符串时，会隐式解码unicode码，所以第二个中文的下标 是 8（utf8中一个中文占3个字节）
	//遍历字符串时，需注意中文索引是按照unicode隐式解码的，下标不连续
	for i, v := range say {
		fmt.Println(i, v)
	}

}
