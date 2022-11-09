package main

import (
	"fmt"
	"runtime"
)

/*
	switch 是编写一连串 if - else 语句的简便方法
	Go 只运行选定的 case，而非之后所有的 case
	Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止
	GO 的switch 的 case 无需为常量，且取值不必为整数
*/
func main() {
	fmt.Print("GO runs on ")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("os x.")
	case "linux":
		fmt.Println("linux.")
	default:
		//windows
		fmt.Println(os)
	}
}
