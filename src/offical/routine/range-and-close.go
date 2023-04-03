package main

import (
	"fmt"
)

//使用channel实现一个斐波那契函数
func fibonacci(len int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < len; i++ {
		ch <- a
		a, b = b, a+b
	}
	/*a, b := 0, 1
	for i := 0; i < len; i++ {
		ch <- a
		a, b = b, a+b
	}*/
	//only  the sender can close channel ,receiver can't close it
	close(ch)
}
func main() {
	ch := make(chan int, 10)
	//attention ： cap() is the size of defined channel ,instantiation  channel invoke len() return 0
	go fibonacci(cap(ch), ch)
	for n := range ch {
		fmt.Println(n)
	}
}
