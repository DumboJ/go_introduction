package main

import "fmt"

func main() {
	var a = [5]int{1, 3, 4, 5, 6}
	fmt.Println(a)
	var s []int = a[1:3]
	fmt.Println(s)
	fmt.Printf("the type of s is %T \n", s)
	fmt.Printf("the type of s is %T", a)
}
