package main

import "fmt"

type Vertex struct {
	a, b float32
}

/*方法只是一个带有函数接收者的函数*/
func vbs(v Vertex) float32 {
	return v.a*v.a + v.b*v.b
}
func main() {
	v := Vertex{2.1, 3.1}
	fmt.Println(vbs(v))
}
