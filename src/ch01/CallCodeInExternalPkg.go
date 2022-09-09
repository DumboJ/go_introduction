package main

import "fmt"
import "rsc.io/quote"

/*调用外部包中代码*/
func main() {
	fmt.Println(quote.Go()) //Don't communicate by sharing memory, share memory by communicating.

	fmt.Println(quote.Hello()) //Hello, world.

	fmt.Println(quote.Opt()) //If a program is too slow, it must have a loop.

}
