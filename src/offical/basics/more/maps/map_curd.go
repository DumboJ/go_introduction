package main

import "fmt"

func main() {
	//init map
	m1 := make(map[string]int)

	//insert
	m1["value"] = 23
	fmt.Println("VLAUE: ", m1["value"])

	//update
	m1["value"] = 22
	fmt.Println("VLAUE: ", m1["value"])

	//delete
	delete(m1, "value")
	fmt.Println("VLAUE: ", m1["value"])

	//通过二值分配确定一个map中是否存在某个值，如果key存在于map中
	/*
		如果key 存在于map中，第二个变量值为 true,反之false
		如果key 不存在于map中，value的值为对应map声明value值的初始值（零值）
	*/
	v, ok := m1["value"]
	fmt.Println("VALUE: ", v, "present? ", ok)
}
