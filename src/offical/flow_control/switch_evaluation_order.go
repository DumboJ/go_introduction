package main

import (
	"fmt"
	"time"
)

/*
	switch 语句的求值顺序
	switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
*/
func main() {
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	fmt.Printf("Today is %s ,time.Saturday value is %s\n", today, time.Saturday+1)
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("Too far away")

	}
}
