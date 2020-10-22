package main

import(
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v, is of type %T\n", book[0], book[0])

	// This is not allowed because strings in Go are immutable
	//book[0] = 116

	// can use slicing on strings
	// slices are 0 based and 1/2 empty (you get 4 but not 11)
	fmt.Println(book[4:11])

	//Print book with a lower case t
	fmt.Println("t" + book[1:])

	// Strings are unicode
	// Multi line string starts with `
	long_string := `
	This is a very
	long string
	over several lines
	...
	`
	fmt.Println(long_string)
}