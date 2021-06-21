package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var pic = make([][]uint8, dx)

	for y:=0; y<dy; y=y+1 {
		pic[y]=make([]uint8, dx)
	}

	for x:=0; x<dx; x+=1 {
		y:=x*x/80
		if y < 0 || y > dy {
			y=0
		}
		pic[x][y]=255
	}
	return pic
}

func main() {
	pic.Show(Pic)
}