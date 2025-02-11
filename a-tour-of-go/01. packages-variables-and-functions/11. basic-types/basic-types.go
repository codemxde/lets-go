package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe  bool
	maxInt int8 = 127
	z complex128 = cmplx.Sqrt(-8 + 5i)
)

func main() {
	fmt.Printf("Type:  %T, Value: %v\n", toBe, toBe)
	fmt.Printf("Type:  %T, Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type:  %T, Value: %v\n", z, z)
}