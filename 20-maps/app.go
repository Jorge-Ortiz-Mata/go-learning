package main

import "fmt"

func main() {
	// Similar to hashes in Ruby.
	// [string] is the datatype for the keys. string is the datatype for values.
	websites := map[string]string{"Google": "google.com", "Amazon": "aws.com"}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Microsoft"] = "azure.com"

	fmt.Println(websites)
	fmt.Println(websites["Microsoft"])

	delete(websites, "Amazon")
	fmt.Println(websites)
}
