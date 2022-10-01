package main

import "fmt"

/*
	GO语言的for循环结构只有一种，即本例中的
			基本for循环由三部分构成
				初始化语句：第一次迭代前执行		通常为短变量声明，该变量声明仅在for的作用域中可见
				条件表达式：每次迭代前求值
				后置语句： 每次迭代的结尾执行

	注意
*/
func main() {
	for i := 3; i < 14; i++ { //for语句后面三个构成部分外没有小括号()，大括号{}是必须的
		fmt.Printf("current value is :%v \n", i)
	}

}
