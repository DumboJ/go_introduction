package main

import "fmt"

type Person struct {
	name string
	age  int
}

//
func (p Person) String() string {
	return fmt.Sprintf("The Person Info :%q,%d\n", p.name, p.age)
}
func main() {
	p := Person{"Tom", 22}
	p1 := Person{"Lily", 23}
	fmt.Println(p, p1)
}
