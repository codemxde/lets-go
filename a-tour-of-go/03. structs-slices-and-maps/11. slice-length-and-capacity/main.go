package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("length: %d cap: %d s: %v\n", len(s), cap(s), s)
}

func main() {
	s := []int {2, 3, 5, 7, 11, 13}

	// * slicing to give 's' zero lenth
	s = s[:0]
	printSlice(s)

	// * Extendings it's length
	s = s[:4]
	printSlice(s)

	// Drop it's first two values
	s = s[2:]
	printSlice(s)
}