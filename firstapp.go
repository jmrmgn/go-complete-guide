package main

import "fmt"

func main() {
	// var greetingText string
	// greetingText = "Hello World!"

	// var greetingText string = "Hello World"

	greetingText := "Hello World!"
	luckyNumber := 17
	evenMoreLuckyNumber := luckyNumber + 5 // 17 + 5

	fmt.Println(greetingText)
	fmt.Println(greetingText)
	fmt.Println(luckyNumber)         // 17
	fmt.Println(evenMoreLuckyNumber) // 22

	// Floats
	var newNumber float64 = float64(luckyNumber) / 3

	fmt.Println(newNumber)

	var defaultFloat float64 = 1.23456789123456789123456
	var smallFloat float32 = 1.23456789123456789123456

	fmt.Println(defaultFloat)
	fmt.Println(smallFloat)
}
