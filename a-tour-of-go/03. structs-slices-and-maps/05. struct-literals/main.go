package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{1, 2}
)

func main() {
	// some random checks
	p.X = 3
	fmt.Println("v1:", v1)
	fmt.Println("p:", p)
	fmt.Println("v2:", v2)
	fmt.Println("v3:", v3)
}