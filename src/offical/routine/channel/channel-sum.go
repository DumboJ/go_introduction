package main

import "fmt"

//利用chan 接收和发送分组数据计算和
func main() {
	ch := make(chan int)
	arr := []int{2, 4, 66, 12, 3, 2}
	//将数组分组后使用协程计算和。并将结果放入channel中保存
	go sumCh(arr[len(arr)/2:], ch)
	go sumCh(arr[:len(arr)/2], ch)

	//取出channel中的值
	a, b := <-ch, <-ch
	fmt.Println(a, b, a+b)
}

func sumCh(arr []int, ch chan int) {
	var sumv int
	for _, v := range arr {
		sumv += v
	}
	//将结果放入channel中
	ch <- sumv
}
