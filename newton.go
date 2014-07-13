package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var y float64
	for i := 0; i < 50; i++ {
		y = z - (((z * z) - x) / (2 * z))
		if math.Abs(z-y) < 0.000000000000001 {
			return z
		}
		z = y
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println()
	fmt.Println(math.Sqrt(2))
}
