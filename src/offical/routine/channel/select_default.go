package main

import (
	"fmt"
	"time"
)

/**
select中的其它分支都没有准备好时，default 分支就会执行
*/
func main() {
	tick := time.Tick(100 * time.Millisecond)
	Boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-Boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
