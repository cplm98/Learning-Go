package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("----")
	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
	}

	//Can use for loop with just condition to imitate while loop
	a := 0
	fmt.Println("----")
	for a < 3 {
		fmt.Println(a)
		a++
	}

	//Can use for loop whithout any condition as a "While True" loop
	b := 0
	fmt.Println("----")
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
}