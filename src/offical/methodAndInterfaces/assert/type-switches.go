package main

import "fmt"

func main() {
	switchType("this is string")
	switchType(true)
	switchType(4.2)
}

func switchType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	case float64:
		fmt.Printf("twice %v = %v \n", v, v*2)
	default:
		fmt.Printf("I don't konw about the type %T \n", v)
	}
}
