package main

import "fmt"

/*todo GO语言中 方法和函数有微小区别 方法中必须要有一个接收者
  todo GO中，函数和包关联，同理，方法可以理解为类型关联的函数
*/

//接收者定义 type receiverName data-type
type UserName string

//方法定义 func (recName receiver) methodName(param type) returnType {}
func (userName UserName) sayHi() {
	fmt.Println("Good Morning, my name is ", userName)
}

//方法的值类型和指针类型
func (userName *UserName) Modify() {
	*userName = UserName("Tom")
}

func main() {
	name := UserName("Lucky")
	name.sayHi()
	name.Modify() //也可以通过 (&name).modify()调用 此时 modify()方法的接受者类型就不用指定指针类型，GO会自动转换
	name.sayHi()
}
