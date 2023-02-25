package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, msg := strconv.Atoi("i")
	if msg != nil {
		fmt.Println(msg) //strconv.Atoi: parsing "i": invalid syntax
	} else {
		fmt.Println(i)
	}
}
