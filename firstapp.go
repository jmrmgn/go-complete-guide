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

	evenMoreLuckyNumber = luckyNumber * 3 // 17 * 3
	fmt.Println(evenMoreLuckyNumber)      // 51

}
