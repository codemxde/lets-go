package main

import (
	"fmt"
)

func Sqrt1(x float64) float64 {
	z := 1.0
	i := 1
	
	for ;i < 10; i++ {
		z -= (z * z - x) / (2 * z)
		fmt.Printf("for i = %d, z = %g\n", i, z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z, diff := float64(1), 1.0
	epsilon := 1e-10
	
	for diff > epsilon {
		prev := z
		z -= (z * z - x) / (2 * z)
		diff = prev - z
	}
	return diff
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println("A dynamic approach:", Sqrt2(2))
}
