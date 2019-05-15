package main

import (
	"golang.org/x/tour/pic"
)

var dx = 5
var dy = 5

func Pic(dx, dy int) [][]uint8 {

	inte := make([]uint8, dx)
	exte := make([][]uint8, dy)

	for j,_ := range exte {
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