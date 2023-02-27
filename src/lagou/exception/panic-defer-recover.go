package main

import "fmt"

//通过 recover()抓取panic终止错误信息

func connectionMySQL(ip, username, password string) {
	if ip == "" {
		panic("ip不能为空")
	}
}
func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	connectionMySQL("", "dumboj", "12453j")
}
