package main

import (
	"fmt"
	"strconv"
)

func main() {
	v, err := strconv.Atoi("_")
	if err != nil {
		//invalid syntax
		fmt.Printf("can't convert to int %v \n ", err)
		return
	}
	fmt.Println("convert to int ,value :", v)
}
