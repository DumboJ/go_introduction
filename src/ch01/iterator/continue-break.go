package main

import "fmt"

func main() {
	fmt.Println(invokeBreak())
	fmt.Println(invokeContinue())
	fmt.Println(invokeWhile())
}

func invokeWhile() int {
	sum := 0
	i := 0
	for i <= 100 {
		sum += i
		i++
	}
	return sum
}

func invokeContinue() int {
	sum := 0
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		sum += i
	}
	return sum
}

func invokeBreak() int {
	sum, i := 0, 0

	for {
		if i > 100 {
			break
		}
		sum += i
		i++
	}
	return sum
}
