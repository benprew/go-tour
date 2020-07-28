package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		img[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			img[y][x] = uint8(x^y)
		}
	}
	return img
}

func main() {
	pic.Show(Pic)
}
