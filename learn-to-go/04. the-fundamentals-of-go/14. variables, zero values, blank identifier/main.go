package main

import (
	"fmt"
)

var package_level string
func main() {
	var age int = 24
	fmt.Printf("He's %d years old\n", age)

	// there's another way apparently
	name := "Keshav"
	fmt.Println("They call him", name)

	var python_type_shit bool = false
	fmt.Printf("You are so %v\n", python_type_shit)

	// look at this 'go' (pun intended)
	a, b, c, d, _, f := 0, "b", 2, "D", 23, "happiness" // not recommended though
	fmt.Println(a, b, c, d, f)

	// another example
	// willItWork := "let's see"
	// var whatAboutThis int = 23

	// * hey look!
	var savage string
	fmt.Println(savage)

	const profession string = "developer"
	// OR
	const hobies = "music"
	
	// fmt.Println(profession, hobies)

	// * grouping declaration
	var (
		playboi string = "carti"
		aubrey = "drake"
		kungfu = "kenny"
		do_it_later int
	)

	fmt.Printf("%s loves %v and so does %s %d", playboi, aubrey, kungfu, do_it_later)

	// * zero value
	var integer_value int
	var string_value string
	fmt.Printf("\nzero value integer: %d", integer_value)
	fmt.Printf("\nzero value string: %s", string_value)

	// string literals
	const song string = `
	Take Me Back To LA
		- Abel Tesfaye
	`
	fmt.Println("fav song right now:", song)

}
