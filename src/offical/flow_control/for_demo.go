package main

import (
	"fmt"
)

/*
	GO语言的for循环结构只有一种，即本例中的
			基本for循环由三部分构成
				初始化语句：第一次迭代前执行		通常为短变量声明，该变量声明仅在for的作用域中可见
				条件表达式：每次迭代前求值
				后置语句： 每次迭代的结尾执行

	注意：
		go 中for语句后可以没有小括号，{}必须
*/
func main() {
	for i := 3; i < 14; i++ { //for语句后面三个构成部分外没有小括号()，大括号{}是必须的
		fmt.Printf("current value is :%v \n", i)
	}
	/*
		for 中初始化语句和后置语句可以省略
		go中 for 语句后的简写后再去掉分号用法相当于 while
	*/
	sum := 1
	for sum < 1000 {
		sum += sum

	}
	fmt.Println(sum)
	/*
		如果省略条件语句，循环则会无限执行
	*/
	for {
		fmt.Println("loading...")
	}
}
