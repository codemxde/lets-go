package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int) int {
	return x + y
}

// revision (additional, not part of tour!)
func addAgain(x, y int) string {
	sum := x + y
	stringified := strconv.Itoa(sum)
	return stringified
}

func main() {
	fmt.Println(add(5, 6))
	fmt.Println("this is fun:", addAgain(5, 6)) // not part of tour
}