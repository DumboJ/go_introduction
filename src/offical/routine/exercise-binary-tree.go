package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/*// Tree /** 树形结构 struct
type Tree struct {
	Left  *Tree
	val   int
	Right *Tree
}*/
// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}
func TestWalk() {
	ch := make(chan int, 10)
	k := 2
	go Walk(tree.New(k), ch)
	for i := 1; i <= 10; i++ {
		val := <-ch
		if val != i*k {
			fmt.Println("Go Walk error ")
			return
		}
	}
	fmt.Println("PASS")
}
func main() {
	TestWalk()
}
