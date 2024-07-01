package main

import (
	"curso-go-intermedio/test/core"
	"fmt"
)

func main() {
	u := core.User{
		Name:  "Valentin Zoia",
		Age:   18,
		Email: "a@a.com",
	}

	isValid, err := core.ValidateUser(u)
	fmt.Println("Is valid:", isValid, "Error:", err)

}
