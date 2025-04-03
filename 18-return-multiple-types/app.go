package main

import "fmt"

func main() {
	res := itReturnsAnything(1)
	fmt.Println(res)

	res = itReturnsAnything(2)
	fmt.Println(res)

	res = itReturnsAnything(3)
	fmt.Println(res)

	other_int := goKnowsTheValueToReturn(1, 2)
	other_string := goKnowsTheValueToReturn("jorge", "ortiz")
	other_float := goKnowsTheValueToReturn(1.2, 4.2)

	fmt.Printf("Values: %v, %v, %v\n", other_int, other_string, other_float)
}

func itReturnsAnything(option int) any {
	switch option {
	case 1:
		return 10
	case 2:
		return "I return string"
	case 3:
		return 23.99
	}

	return nil
}

func goKnowsTheValueToReturn[T int | string | float64](a, b T) T {
	return a + b
}
