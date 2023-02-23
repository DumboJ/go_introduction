package main

import "fmt"

func main() {
	var a = [5]int{1, 3, 4, 5, 6}
	fmt.Println(a)
	var s []int = a[1:3]
	fmt.Println(s)
	fmt.Printf("the type of s is %T \n", s)
	fmt.Printf("the type of a is %T \n", a)
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	slice1 := append(slice, 3, 2)
	fmt.Println(slice1)

	for _, v := range a {
		fmt.Println(v)
	}
}
