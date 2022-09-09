package main

import "fmt"

func main() {
	f := fib()
	fmt.Println(f(), //a=1,b=1
		f(), //a=1,b=2
		f(), //a=2,b=3
		f(), //a=3,b=5
		f()) //a=5,b=8
}
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}

}
