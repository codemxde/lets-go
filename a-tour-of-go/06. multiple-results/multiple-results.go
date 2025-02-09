package main

import "fmt"

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

// revision
func swoppyTroppy(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println(swap("hello", "world"))

	// revision
	x, y := 10, 20
	x, y = swoppyTroppy(x, y)
	fmt.Printf("x: %d and y: %d", x, y)
}