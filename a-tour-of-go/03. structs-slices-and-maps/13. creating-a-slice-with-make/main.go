package main

import "fmt"

func printSlice(str string, s []int)  {
	fmt.Printf("%s: length: %d cap: %d %v\n",str, len(s), cap(s), s)
}

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
	// fmt.Println(b == nil)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}