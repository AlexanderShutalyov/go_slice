package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	exte := make([][]uint8, dy)

	for j,_ := range exte {
		inte := make([]uint8, dx)
		for i,_ := range inte {
			inte[i] = uint8(i*j)
		}
		exte[j] = inte
	}

	return exte

}

func main() {
	pic.Show(Pic)
}