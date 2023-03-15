package main

import "fmt"

//GO中接口都是隐式声明的
//defer an Interface
type I interface {
	M()
}

//defer an Struct
type T struct {
	S string
}

//make the T implements interface I ,we can only make the Type as the func receiver,otherwise the T has same func with interface
func (t T) M() {
	fmt.Println(t.S)
}
func main() {
	var i I = T{"implicitly defined interface"}
	i.M()
}
