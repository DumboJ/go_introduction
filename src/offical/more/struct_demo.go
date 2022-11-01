package main

import "fmt"

type coordinate_axis struct {
	X, Y int
}

//struct 字面意义：struct 可以通过 name:value 来对其中元素进行赋值
//& return a more to struct value
var (
	a  = coordinate_axis{1, 2}
	b  = coordinate_axis{Y: 3}
	c  = coordinate_axis{} //有对应struct结构体的init value
	p1 = &coordinate_axis{4, 5}
)

func main() {
	//struct 赋值
	fmt.Println(coordinate_axis{1, 2})

	//struct field are accessed using a dot
	p := coordinate_axis{12, 3}
	println(p.X)

	//more to struct
	v := coordinate_axis{1, 2}
	t := &v
	t.X = 1e9 //(*t).X
	fmt.Println(v)

	fmt.Println(a, b, c, p1)
}
