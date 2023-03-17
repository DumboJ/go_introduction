package main

import (
	"fmt"
	"strconv"
)

type IPAddr struct {
	name string
	ip   [4]byte
}

func (i IPAddr) String() string {
	lastIndex := len(i.ip) - 1
	var res string = i.name + ("--")
	for index, v := range i.ip {
		res += strconv.Itoa(int(v))
		if index < lastIndex {
			res += "."
		}
	}
	return res
}
func main() {
	i := IPAddr{"localhost", [...]byte{127, 0, 0, 1}}
	fmt.Println(i)
}
