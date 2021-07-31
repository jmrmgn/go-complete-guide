package main

import (
	"fmt"

	"github.com/yourorg/firstapp/greeting"
)

func oldMain() {
	// var greetingText string
	// greetingText = "Hello World!"

	// var greetingText string = "Hello World"

	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5 // 17 + 5

	fmt.Println(greeting.GreetingText)
	fmt.Println(greeting.GreetingText)
	fmt.Println(luckyNumber)         // 17
	fmt.Println(evenMoreLuckyNumber) // 22

	// Floats
	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	var defaultFloat float64 = 1.23456789123456789123456
	var smallFloat float32 = 1.23456789123456789123456

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)

	// Rune
	var firstRune rune = '$'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	// Byte
	var firstByte byte = 'a'
	fmt.Println(firstByte)

	// Concatenating Strings
	firstName := "John"
	lastName := "Doe"
	fullName := fmt.Sprintf("%v %v", firstName, lastName)
	age := 20
	greetingSentence := fmt.Sprintf("Hi, I am %v and I am %v (Type: %T) years old.", fullName, age, age)

	fmt.Println(greetingSentence)
	// fmt.Println("9" + 1)

	// Multiline
	// multiline := `This is a multiline string!
	// It is quite long.

	// And spans multiple lines`
	// fmt.Println(multiline)

	fmt.Println("PI from pi.go file", pi)
}
