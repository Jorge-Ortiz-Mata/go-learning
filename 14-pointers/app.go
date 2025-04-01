package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("Age:", *agePointer) // Get the value store in memory
	fmt.Println("Age:", agePointer)  // Get the address value store in memory

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult years:", adultYears)

	mutateAgeInMemory(agePointer)
	fmt.Println("Age:", age) // Get the value store in memory
}

func getAdultYears(age *int) int {
	return *age - 18
}

func mutateAgeInMemory(age *int) {
	*age = 10
}
