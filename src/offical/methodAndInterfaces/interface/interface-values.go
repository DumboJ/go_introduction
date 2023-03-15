package main

import "fmt"

//interface 可以作为参数和返回值 可以看成是（value,type）形式，接口值保存具体类型的值和类型，调用时执行的是底层类型的同名方法
type IT interface {
	M()
}
type F float64
type user struct {
	name, alias string
}

func (p F) M() {
	fmt.Println(p, "\n")
}
func (u user) M() {
	fmt.Println(u.name + "--" + u.alias)
}

//调用时，接口保存的底层具体类型的值
func describe(i IT) {
	fmt.Printf("I vlues : %v and type :%T\n", i, i)
}
func main() {
	var i IT
	i = user{"LILI", "LI"}
	i.M()
	describe(i)

	i = F(6.23)
	i.M()
	describe(i)
}
