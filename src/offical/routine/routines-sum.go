package main

import "fmt"

func main() {
	//多线程计算数组的和
	ch := make(chan int)
	arr := []int{3, 4, 22, 5, 19, 2}
	//将数组切分放进channel中并计算
	go sum(arr[:len(arr)/2], ch)
	go sum(arr[len(arr)/2:], ch)
	x, y := <-ch, <-ch

	fmt.Println(x, y, x+y)
}

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}
