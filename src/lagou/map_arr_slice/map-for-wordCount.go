package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	str "strings"
)

//统计一句话中每个单词出现的次数
func wordCount(s string) map[string]int {
	arr := str.Fields(s)
	m := make(map[string]int)
	for _, val := range arr {
		m[val]++
	}
	return m
}
func main() {
	v := wordCount("I eat a donut. Then I eat another donut.")
	fmt.Println(v)
	wc.Test(wordCount)
}
