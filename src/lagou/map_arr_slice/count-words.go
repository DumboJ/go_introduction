package main

import (
	"fmt"
	"strings"
)

func countWords(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	return m
}
func main() {
	res := countWords("this is an little pag , how much will you speed pag")
	fmt.Println(res)
}
