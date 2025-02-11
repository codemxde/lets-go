package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum;
	}
	fmt.Println(sum)

	for i:= 1; i < 10; {
		i += i*2
		fmt.Println("i:", i)
	}
}