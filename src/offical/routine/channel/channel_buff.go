package main

import "fmt"

/**
*带缓冲的信道（初始化包含信道容量）

 */
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1
	ch <- 1
	fmt.Println(<-ch)
	//fatal error: all goroutines are asleep - deadlock! 超出信道容量cause error
	fmt.Println(<-ch)
}
