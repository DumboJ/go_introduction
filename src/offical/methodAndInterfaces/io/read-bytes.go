package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "this is string to read."
	b := make([]byte, 8)
	r := strings.NewReader(str)
	for {
		n, err := r.Read(b)
		// type %v --> byte
		fmt.Printf("n := %v ,err := %v ,b = %v \n", n, err, b)
		//type %s --> string
		fmt.Printf("b[:n] = %s \n", b[:n])
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
