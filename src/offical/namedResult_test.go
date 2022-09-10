package offical

import "testing"

func Test_returnNamed(t *testing.T) {
	t.Log(split(10))
	t.Log(appendStr("string test .. "))
}

/*命名返回值
Go的返回值可以被命名，它们被视作定义在函数顶部的变量
没有参数的 return 语句返回已命名的返回值~直接返回
直接返回应当用在短函数中，长函数中影响代码可读性
*/
func split(sum int) (x, y int) {
	x = sum*7 + 10
	y = sum + 5
	return
}

func appendStr(start string) (d, f string) {
	d = start + "append me test"
	f = start + "append test end"
	return
}
