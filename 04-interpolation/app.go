package main

import "fmt"

func main() {
	const name string = "Jorge Ortiz"
	const age int = 24

	// %v is a format specifier that formats the value in a default format
	fmt.Printf("Hey, my name is %v and I am %v years old\n", name, age)
}
