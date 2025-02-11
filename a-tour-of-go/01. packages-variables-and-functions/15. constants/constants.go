package main

import "fmt"

const global = "global"

func main() {
	const hello = "world"
	const truth = true

	fmt.Printf("hello %q\nis go %s ? %v", hello, global, truth)
}