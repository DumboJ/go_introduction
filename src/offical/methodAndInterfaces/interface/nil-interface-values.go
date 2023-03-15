package main

import "fmt"

//空的接口不保存具体值也不保存具体类型，attention :todo!! when interface{} ,can't spark  panic exception.only the (value<nil>,type<nil>)
type ni interface {
	M()
}

func main() {
	var i ni
	//spark panic exception : panic: runtime error: invalid memory address or nil pointer dereference
	desci(i)
	i.M()
}
func desci(i ni) {
	fmt.Printf("nil interface value : %v , nil interface type :%T\n", i, i)
}
