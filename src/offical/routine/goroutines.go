package main

import (
	"fmt"
	"time"
)

func main() {
	//协程是Go管理的运行时轻量级线程
	go say("this is go ")
	say("routines")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(300 * time.Millisecond)
	}
}
