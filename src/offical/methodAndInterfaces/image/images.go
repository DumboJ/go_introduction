package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MyPic struct {
	Width, Height int
	colr          uint8
}

// ColorModel MyPic 实现image接口三个方法
//1.颜色模式
func (pic *MyPic) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds 2.图片边界
func (pic *MyPic) Bounds() image.Rectangle {
	return image.Rect(0, 0, pic.Width, pic.Height)
}

// At 3.生成图像中某个点的方法
func (pic *MyPic) At(x, y int) color.Color {
	return color.RGBA{R: uint8(x) + pic.colr, G: uint8(y) + pic.colr, B: uint8(255), A: uint8(255)}
}
func main() {
	m := MyPic{100, 100, 120}
	pic.ShowImage(&m)
}
