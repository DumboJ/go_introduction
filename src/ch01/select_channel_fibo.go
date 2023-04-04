package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	select_fibo(ch, quit)
}

func select_fibo(ch chan int, quit chan int) {
	m, n := 0, 1
	for {
		select {
		case ch <- m:
			m, n = n, m+n
		case <-quit:
			fmt.Println("program quit.")
			return
		}
	}
}
