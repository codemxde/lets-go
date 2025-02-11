package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 5, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)

	fmt.Printf("f: %v, z: %v", f, z)
}