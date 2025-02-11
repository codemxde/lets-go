package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum/2;
	y = sum/3;

	return
}

// revsion
func nakedReturn(sum int) (x, y int) { // here x, y are initialized with their 'zero' value
	x += 2
	y -= 3
	return x, y
}

func main() {
	fmt.Println(split(15))

	// revsion
	x, y := nakedReturn(10)
	fmt.Println("naked return:", x, y)
}