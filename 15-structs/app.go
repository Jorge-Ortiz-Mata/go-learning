package main

import (
	"example/15-structs/user"
	"fmt"
)

func main() {
	u, err := user.New("Jorge Ortiz", "jorge@gmail.com", "Pass12")
	uAdmin := user.NewAdmin("Jorge Ortiz", "jorge@gmail.com", "Pass12", "gh_token1234")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.Password)
		u.UpdatePassword()
		fmt.Println(u.Password)
	}

	// Inheritance
	fmt.Println(uAdmin)
	fmt.Println(uAdmin.User)
	fmt.Println(uAdmin.GhToken)
	uAdmin.ShowDetails()
}
