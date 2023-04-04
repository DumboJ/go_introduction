package main

import "fmt"

/*普通实现fibonacci数列*/
func main() {
	ints := make(chan int, 10)
	go simple_fibo(cap(ints), ints)
	for i := range ints {
		fmt.Println(i)
	}
}

func simple_fibo(len int, ch chan int) {
	m, n := 0, 1
	for i := 0; i < len; i++ {
		ch <- m
		m, n = n, m+n
	}
	close(ch)
}
