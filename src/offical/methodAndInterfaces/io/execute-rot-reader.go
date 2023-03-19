package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

//实现rot13Reader的Read方法
func (r rot13Reader) Read(b []byte) (int, error) {
	t := make([]byte, 20)
	i, err := r.r.Read(t)
	for index, value := range t[:i] {
		b[index] = rot13(value)
	}
	return i, err
}
func rot13(b byte) byte {
	switch {
	case b >= 'A' && b <= 'M' || b >= 'a' && b <= 'm':
		b += 13
	case b >= 'N' && b <= 'Z' || b >= 'n' && b <= 'z':
		b -= 13
	}
	return b
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
