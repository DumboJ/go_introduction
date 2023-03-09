package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   = 0
	mutex sync.Mutex
)

// concurrent wrong
//resolve : add sync.Mutex
func main() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	//防止程序提前退出
	time.Sleep(2 * time.Second)
	fmt.Println(sum) //990
}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
