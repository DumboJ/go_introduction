package main

import "fmt"

func main() {
	ch := make(chan int)
	go count(ch, 10)
	go count(ch, 5)
	s, c := <-ch, <-ch
	fmt.Println(s, c, s+c)
}
func count(ch chan int, step int) {
	sum := 0
	for i := 0; i < step; i++ {
		sum += i
	}
	ch <- sum
}
