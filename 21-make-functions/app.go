package main

import "fmt"

type stringMap map[string]string

func main() {
	userNames := make([]string, 3, 3)
	userNames[0] = "Luis"
	userNames[1] = "Zaira"
	userNames[2] = "Andres"
	userNames = append(userNames, "Jorge")
	userNames = append(userNames, "Ana")
	userNames = append(userNames, "Maria")
	userNames = append(userNames, "Mario")
	for _, value := range userNames {
		fmt.Println("Username Value")
		fmt.Println(value)
	}

	languages := make(stringMap, 2)

	languages["ENG"] = "english"
	languages["SPA"] = "spanish"

	for _, value := range languages {
		fmt.Println("Languages Value")
		fmt.Println(value)
	}
}
