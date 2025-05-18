package main

import (
	"fmt"
	"math"
)

func main() {
	x := []float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for _, v := range x {
		arcsin := math.Asin(v)
		arccos := math.Acos(v)
		y := math.Pow(arcsin*arcsin+math.Pow(arccos, 4), 3)
		fmt.Println("При x =", v, "y =", y)
	}
}
