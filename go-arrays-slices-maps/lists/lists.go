package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99) // push the value in prices
	prices = prices[1:]           // removing the 1st value of prices[]
	fmt.Println(prices)

	discountPrices := []float64{12.99, 9.50, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// origPrices := prices[1:3]   // [9.99, 45.99] include 1 and exclude the 3 => 1~2
// 	// otherPrices := prices[:3]       //[10.99 9.99 45.99] from the beginning to the given index(3)
// 	featuredPrices := prices[1:] // [9.99 45.99 20] from index 1 to the end of array
// 	featuredPrices[0] = 199.99   // will change the featuredPrices[0]
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)

// 	// Getting the length of the array or other
// 	fmt.Println(len(highlightedPrices))

// 	// Getting the current capacity of that array, not the original
// 	fmt.Println(cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3] // can reslice
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
