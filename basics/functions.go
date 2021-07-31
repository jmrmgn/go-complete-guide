package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := add(1, 5)
	diff := minus(10, 4)

	result := fmt.Sprintf("The sum is %v", sum)
	fmt.Println(result)
	printNumber(sum)

	randomNum1, randomNum2 := generateRandomNumbers()
	fmt.Println(randomNum1, randomNum2)
	fmt.Printf("The difference is %v", diff)
}

// Named Return Values: Automatic return because sum is defined as the return value
func generateRandomNumbers() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)
	return
}

// Named Return Values: Automatic return because sum is defined as the return value
func add(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return
}

// Vanilla function
func minus(minued int, subtrahend int) int {
	diff := minued - subtrahend
	return diff
}

func printNumber(num int) {
	fmt.Println(num)
}
