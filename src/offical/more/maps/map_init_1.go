package main

import "fmt"

//位置点struct
type Coord struct {
	lat, lon float32
}

var s = map[int]Coord{
	//模拟北京市位置坐标
	//如果value的顶级类型只是一个类型名，可以省略,类似于Coord
	110000: Coord{151.000384, 41.33334},
	//110000: {151.000384, 41.33334},

	410000: {171.000384, 38.21548},
}

func main() {
	/*map 类似于struct ,但map必须有键值，value 可以为nil */
	fmt.Println(s)
}
