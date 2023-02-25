package main

import "fmt"

//定义接口Animal ,接口中有接口方法
type Animal interface {
	//接口方法
	Say() string
}

//定义一个具体动物
type Cat struct {
	name     string
	category string
	voice    string
}

//接口的实现
func (cat Cat) Say() string {
	return fmt.Sprintf("the cat of %s voice is %s", cat.name, cat.voice)
}

func main() {
	//既然Cat 结构体类型 实现了接口Animal 的Say 方法 就可以通过Cat实例 调用
	cat := Cat{"dudu", "波斯猫", "miaomiao"}
	sendVoice(cat)
}

// sendVoice方法 接收了 Animal类型的参数
func sendVoice(animal Animal) {
	fmt.Println(animal.Say())
}
