package main

import (
	"fmt"
	"math"
)

// Initial implementation
func EquableTriangle(a, b, c int) bool {
	var area, perimeter float64
	aa := float64(a)
	bb := float64(b)
	cc := float64(c)
	perimeter = aa + bb + cc
	sPerimeter := 0.5 * perimeter
	area = math.Sqrt(sPerimeter * (sPerimeter - aa) * (sPerimeter - bb) * (sPerimeter - cc))
	return area == perimeter
}

// sides can be extracted as separate float64 vars, or can be converted in the expression
// no need to define perimeter or semiPerimeter in the beginning
// no need to separate area as a separate var (depends of course)
func NicerEquableTriangle(a, b, c int) bool {
	perimeter := float64(a + b + c)
	sp := 0.5 * perimeter
	return perimeter == math.Sqrt(sp*(sp-float64(a))*(sp-float64(b))*(sp-float64(c)))
}

func RunEquableTriangle() {
	fmt.Println(NicerEquableTriangle(5, 12, 13))
}

// func main() {
// 	RunEquableTriangle()
// }
