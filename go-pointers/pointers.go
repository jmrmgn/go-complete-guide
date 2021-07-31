package main

import "fmt"

/*
	& - declaring a pointer => it will return the address
	* - deferencing a pointer => it will return the value from the given address
*/
func main() {
	age := 32
	
	fmt.Println(age)

	myAge := &age
	// other way to assign a pointer
	// var myAge *int
	// myAge = &age
	fmt.Println(*myAge) // dereferencing

	*myAge = 33 // changing the value of myAge by adding *

	fmt.Println(*myAge)
	fmt.Println(age)

	doubledAge := double(myAge)
	fmt.Println(doubledAge)
	fmt.Println(age)
}

func double(number *int) int {
	result := *number * 2
	*number = 100 // this will change the value from the given address
	return result
}