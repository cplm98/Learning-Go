
package main

import (
	"fmt"
)

func main() {
	// var x float64
	// var y float64

	// x = 1
	// y = 2
	// ^ can be replaced by ( := lets the variable infer its type)
	x, y := 1.0, 2.0

	fmt.Printf("x=%v, type of x is %T\n", x, x)
	fmt.Printf("y=%v, type of y is %T\n", y, y)

	// var mean float64
	mean := (x + y) / 2.0
	// type system is very tight
	fmt.Printf("Mean is %v\n", mean)
}