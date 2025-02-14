package main

import (
	"fmt"
	"math"
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
	i := 0
	
	for diff > epsilon {
		prev := z
		z -= (z * z - x) / (2 * z)
		diff = math.Abs(z - prev)
		i++
	}
	fmt.Println("number of iterations:", i)
	return z
}

func Sqrt3(x float64) float64 {
	z := x/6
	diff := z
	i := 1

	for diff > 0 {
		prev := z
		z -= (z * z - x) / (2 * z)
		diff = math.Abs(z - prev)
		i++
	}
	fmt.Println("no of iterations:", i)
	return z
}

func main() {
	fmt.Println(Sqrt1(10), "\n")
	fmt.Println("A dynamic approach:", Sqrt2(10), "\n")
	fmt.Println("Modified Approach:", Sqrt3(10))
}
