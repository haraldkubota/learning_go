package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	width, height int
	color uint8
}

func (p Image) ColorModel () color.Model {
  return color.RGBAModel
}

func (p Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, p.width, p.height)
}

func (p Image) At (x, y int) color.Color {
  return color.RGBA{p.color, p.color, 255, 255}
  // return color.RGBA{p.color + uint8(x), p.color + uint8(y), 128, 255}
}

func main() {
	m := Image{250,250,0}
	pic.ShowImage(m)
}
