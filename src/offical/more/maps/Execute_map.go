package main

import (
	"fmt"
	"strings"
)

/*
	map训练，统计一句话中word的出现次数，wordCount
*/
func WordCount(s string) map[string]int {
	r := make(map[string]int)
	arr := strings.Fields(s)

	for _, val := range arr {
		//strings.Fields处理后的数组元素 val 作为 map的key，累加 作为map的value
		r[val]++
	}
	return r
}
func main() {
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
