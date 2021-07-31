package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := taskOne()

	taskTwo(hobbies)

	mainHobbies := taskThree(hobbies)

	taskFour(mainHobbies)

	fmt.Println(taskFive())

	products := taskSix()
	fmt.Println(products)
}

func taskOne() (hobbies [3]string) {
	hobbies = [3]string{"Eat", "Sleep", "Code"}
	fmt.Println(hobbies)

	return
}

func taskTwo(hobbies [3]string) {
	firstHobby := fmt.Sprintf("The first element (standalone): %v", hobbies[0])
	fmt.Println(firstHobby)

	selectedHobbies := fmt.Sprintf("The second and third element combined as a new list: %v", hobbies[1:])
	fmt.Println(selectedHobbies)
}

func taskThree(hobbies [3]string) []string {
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	return mainHobbies
}

func taskFour(mainHobbies []string) {
	fmt.Println("Task Four")
	fmt.Println(mainHobbies)
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
}

func taskFive() []string {
	goals := []string{"Learn", "Earn"}
	goals[1] = "Earning"
	goals = append(goals, "Enjoy")

	return goals
}

func taskSix() []Product {
	products := []Product{
		{id: "a-book", title: "A book", price: 12.99},
		{id: "ball-fondlers", title: "Ball Fondlers", price: 6.99},
	}

	newProduct := Product{id: "carpet", title: "Carpet", price: 75.99}
	products = append(products, newProduct)

	return products
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
