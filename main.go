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
	for i := 0.26; i <= 0.66; i += 0.08 {
		arcsin := math.Asin(i)
		arccos := math.Acos(i)
		y := math.Pow(arcsin*arcsin+math.Pow(arccos, 4), 3)
		fmt.Println("При x =", i, "y =", y)
	}
}
