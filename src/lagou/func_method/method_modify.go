package main

import "fmt"

//接收者类型
type receiver struct {
	name    string
	context string
}

//定义方法
func (r receiver) show() {
	fmt.Println("receiver name is ", r.name)
	fmt.Println("receiver context is ", r.context)
}
func main() {
	//初始化
	rv := receiver{"first", "receiver"}
	//调用show 方法
	rv.show()
	//通过指针调用modify 改变结构体的值
	cg := receiver{"three", "multi"}
	(&rv).modify(&cg)
	//另一种调用形式，GO会自动转换
	rv.modify(&cg)
	//改变内容后调用show方法
	rv.show()
}
func (r *receiver) modify(r2 *receiver) {
	*r = receiver{r2.name, r2.context}
}
