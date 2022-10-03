package main

import "time"

/*
	没有条件的 switch 同 switch true 一样。

	这种形式能将一长串 if-then-else 写得更加清晰。
*/
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("Good morning")
	case t.Hour() < 17:
		println("Good afternoon")
	default:
		println("Good evening")
	}
}
