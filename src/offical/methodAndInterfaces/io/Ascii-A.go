package main

import "golang.org/x/tour/reader"

type MyReader struct {
}

//实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流
//Read方法用字节填充给定的字节切片并返回填充的字节数和错误值，
//数据流结尾时返回 io.EOF 的错误
func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}
func main() {
	reader.Validate(MyReader{})
}
