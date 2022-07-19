package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

func generateColors(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			//pic[y][x] = (uint8(x) + uint8(y)) / 2
			pic[y][x] = uint8(x) * uint8(y)
		}
	}

	return pic
}

func newImage(w, h int) Image {
	return Image{w, h, generateColors(w, h)}
}

type Image struct {
	w, h   int
	colors [][]uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.colors[x][y], i.colors[x][y], 255, 255}
}

func main() {
	m := newImage(300, 300)
	pic.ShowImage(m)
}
