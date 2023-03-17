package main

import "fmt"

//type assert ->> variable.(type)
func main() {
	var i interface{} = "type assert"
	//print value
	fmt.Println(i.(string))

	//same type return type with the value
	s, ok := i.(string)
	fmt.Printf("value = %v ,same type : %t\n", s, ok)

	//spark panic : interface conversion : interface{} is string ,not float64
	/*t := i.(float64)
	fmt.Println(t)*/

	//no panic receive boolean value ,not same type ,value is the type zero value
	t, isS := i.(float64)
	fmt.Printf("value = %v ,same type : %t \n", t, isS)
}
