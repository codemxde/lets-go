package main

import "fmt"

func addTwoNumbers(x, y int) int {
	return x + y
}

func printAndSum(x, y int, str string) int {
	fmt.Println(str)
	return x + y
}

func main() {
	fmt.Println("Sum is", addTwoNumbers(12, 22))
	fmt.Println(printAndSum(12, 23, "take me back to la!"))
}