package main

import (
	"errors"
	"fmt"
)

//GO中提供包裹错误信息。也可以解开获取原始错误信息
func main() {
	//包裹
	e := errors.New("原始错误e")
	m := fmt.Errorf("包裹了一个错误 %w", e)
	fmt.Println(m)

	//解开
	fmt.Println(errors.Unwrap(m))

	//os.ErrExist  如果多层嵌套则两个error不相等 ，判断方法 使用 errors.Is(err, target error)
	fmt.Println(errors.Is(m, e))
}
