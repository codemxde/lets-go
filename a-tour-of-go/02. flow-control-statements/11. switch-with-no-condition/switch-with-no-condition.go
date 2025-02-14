package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Hour()

	switch {
	case t < 12:
		fmt.Println("Good Morning!")
	
	case t < 17:
	fmt.Println("Good Afternoon")

	default:
		fmt.Println("Good Evening")
	}
}