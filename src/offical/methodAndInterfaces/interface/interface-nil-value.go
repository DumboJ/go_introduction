package main

import "fmt"

//interface
type ifc interface {
	M()
}
type TP struct {
	S string
}

func (t *TP) M() {
	//when receiver is nil
	if t == nil {
		fmt.Println("<nil vlaue>")
		return
	}
	fmt.Println(t)
}
func main() {
	//defined ifc function variable
	var i ifc
	//defined nil pointer of struct
	var t *TP
	//receive
	i = t
	//I vlues : <nil> and type :*main.TP
	desc(i)
	t.M()

}

//保存了具体值为nil的接口自身并不为nil
func desc(i ifc) {
	fmt.Printf("I vlues : %v and type :%T\n", i, i)
}
