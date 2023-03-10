package main

import "fmt"

type author struct {
	name   string
	branch string
	salary float32
}

func (a author) display() {
	fmt.Println(a.name)
	fmt.Println(a.branch)
	fmt.Printf("the slary is %g \n", a.salary)
}
func main() {
	a := author{
		"lucy",
		"chinese paper",
		28660.30,
	}
	a.display()
}
