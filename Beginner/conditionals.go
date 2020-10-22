package main

import(
	"fmt"
)

func main() {
	x := 10 

	if x > 5 {
		fmt.Println("X is Big")
	}

	if x > 100 {
		fmt.Println("X is very Big")
	} else {
		fmt.Println("but not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x < 20 || x > 20 {
		fmt.Println("x is out of range")
	}

	//if can have an optional initialization statement
	a := 11.0
	b := 20.0
	// if fraction -> then comparison
	if frac := a / b; frac > .5 {
		fmt.Println("a is more than half of b")
	}

	// Switch statements
	// don't have to specify break, case doesn't have to be a number
	y := 3
	switch y {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("many")
	}

	// Can run empty switch, where case statement will have the condition
	z := 50
	switch {
	case z > 100:
		fmt.Println("Z is very big")
	case z < 10:
		fmt.Println("Z is very small")
	default:
		fmt.Println("Z is just right")
	}
}