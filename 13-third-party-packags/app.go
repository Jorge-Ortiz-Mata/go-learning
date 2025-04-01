package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	fmt.Println("Hello Third Party Pakages")
	var username string = gofakeit.Name()

	fmt.Println(username)
}
