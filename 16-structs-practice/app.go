package main

import "example/practice/book"

func main() {
	// NOK example - without title
	bookOne := book.New("", "This is a brief content of it")
	bookOne.Call()

	// NOK example - without content
	bookTwo := book.New("Harry Potter", "")
	bookTwo.Call()

	// NOK example - content with length less than 20
	bookThree := book.New("Harry Potter", "Not enough")
	bookThree.Call()

	// OK example
	bookFour := book.New("Harry Potter", "This is a brief content of the Harry Potter book")
	bookFour.Call()
}
