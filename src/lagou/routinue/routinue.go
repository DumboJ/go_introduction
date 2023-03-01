package main

import (
	"fmt"
	"time"
)

func main() {
	//新建go的一个协程，该协程不阻塞main协程执行
	go fmt.Println("this is a go thread")
	//主协程
	fmt.Println("this is main thread")
	//main 协程等待一秒，否则main routine 执行完成后程序就退出了，看不到 go关键开启的协程打印结果
	time.Sleep(time.Second)
}
