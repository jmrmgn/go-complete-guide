package main

import (
	"fmt"
)

func main() {
	// Name
	var firstName string = "John"
	lastName := "Doe"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Age
	var currentYear int = 2021
	birthYear := 1996
	age := currentYear - birthYear

	fmt.Println(age)

	// Overriding the currentYear
	currentYear += 1
	age = currentYear - birthYear

	// New Age
	fmt.Println(age)
}
