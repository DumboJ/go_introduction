package main

import (
	"errors"
	"log"
)

func main() {
	//todo 1. 函数定义
	definedFunc()

	//todo 2.匿名函数 + 闭包
	noNameFunc_closePkg()
}

func noNameFunc_closePkg() {
	//匿名函数展示
	sum := func(a, b int) int {
		return a + b
	}
	log.Println(sum(3, 5))

	//基于匿名函数，可以在已声明的函数中再定义函数（函数嵌套） 匿名函数可以使用所在函数定义的变量等
	//incr_sum 相当于一个 赋值为 匿名函数类型的变量 不是函数名字 incr_sum -- func()
	incr_sum := increme() //此时返回：func() 1，这一步并没有调用匿名函数内部i++的操作，只是返回i的值

	log.Println(incr_sum()) //此时调用的是内部匿名函数func() 返回：经过上一步，此时increme()中 i = 1,
	// 由于Go闭包的能力，匿名函数持有外部函数的变量i。上一步调用后后返回1,这步调用进匿名函数内部，执行i++
	log.Println(incr_sum())
	log.Println(incr_sum())
	log.Println(incr_sum())
}

func increme() func() int {
	i := 1
	//相当于 return func() int{ return 1} 匿名方法返回类型 int -- 1
	return func() int {
		i++
		return i
	}
}

func definedFunc() {
	//1.1 单值返回
	a := definedFunc_oneReslut(2, 4)
	log.Println(a)

	//1.2 多值返回
	b, err := definedFunc_multiResult(3, 2)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(b)
	}

	//1.3 命名返回参数
	sum, error := definedFunc_deferdResult(3, 2)
	log.Println(sum, error)

	//1.4 可变参数 -- todo 多个参数类型时，可变参数放最后
	sumMulti := sum_varargues(1, 2, 3, 5)
	log.Println(sumMulti)
}

func sum_varargues(params ...int) int {
	sum := 0
	for _, v := range params {
		sum += v
	}
	return sum
}

func definedFunc_deferdResult(i int, i2 int) (sum int, err error) {
	if i < 0 || i2 < 0 {
		return 0, errors.New("参数不能为负数！")
	}
	sum = i + i2
	err = nil
	return
}

func definedFunc_multiResult(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("参数应为 > 0 的数值")
	}
	return a + b, nil
}

//定义  func funcName(params) result{ body }
func definedFunc_oneReslut(a, b int) int {
	return a + b
}
