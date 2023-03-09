package main

import "fmt"

var sum = 0

// concurrent wrong
func main() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	fmt.Println(sum) //990
}

func add(i int) {
	sum += i
}
