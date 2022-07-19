package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pic[y][x] = (uint8(x) + uint8(y)) / 2
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
