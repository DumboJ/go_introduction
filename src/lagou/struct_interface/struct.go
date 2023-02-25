package main

import (
	"fmt"
)

type Address struct {
	provence string
	city     string
}

type User struct {
	name string
	age  uint
	addr Address
}

func main() {
	//defined struct
	var user User
	user.age = 14
	user.name = "Lock"
	addr := Address{"云南", "弥勒"}
	user.addr = addr
	fmt.Println(user)

	//字面量初始化
	u := User{"Tony", 22, Address{"云南", "昆明"}}
	//指定字段名式 字面量初始化
	ub := User{"Bob", 22, Address{provence: "云南"}}

	fmt.Println(u)
	fmt.Println(ub)

	//访问结构体中字段 with dot
	println(u.name, u.addr.provence)

	//Person和Address都实现了fmt.Stringer的String()方法
	printString(u)
	printString(u.addr)

}

//让User和Address都实现 fmt.Stringer的String() string 方法
func (user User) String() string {
	return fmt.Sprintf("the user of %s age is %d", user.name, user.age)
}

func (addr Address) String() string {
	return fmt.Sprintf("address is %s and city is %s", addr.provence, addr.city)
}

//打印fmt.Stringer接口的函数
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}
