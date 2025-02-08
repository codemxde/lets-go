package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum/2;
	y = sum/3;

	return
}

func main() {
	fmt.Println(split(15))
}