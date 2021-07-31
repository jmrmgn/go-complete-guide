package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	id        int
	title     string
	shortDesc string
	price     float64
}

func (book *Book) store() {
	file, _ := os.Create(fmt.Sprintf("%v-%v.txt", book.id, book.title))

	content := fmt.Sprintf(`
		ID: %v
		Title: %v
		Description: %v
		Price: USD %.2f
	`, book.id, book.title, book.shortDesc, book.price)

	file.WriteString(content)

	file.Close()
}

func (book *Book) outputData() {
	fmt.Printf("ID: %v\n", book.id)
	fmt.Printf("Title: %v\n", book.title)
	fmt.Printf("Description: %v\n", book.shortDesc)
	fmt.Printf("Price: USD %.2f\n", book.price)
}

func NewBook(id int, title string, shortDesc string, price float64) *Book {
	return &Book{
		id:        id,
		title:     title,
		shortDesc: shortDesc,
		price:     price,
	}
}

func main() {

	createdBook := getBookInfo()

	createdBook.outputData()
	createdBook.store()
}

func getBookInfo() *Book {
	fmt.Println("Enter the book data.")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Book ID: ")
	titleInput := readUserInput(reader, "Book Title: ")
	descInput := readUserInput(reader, "Book Description: ")
	priceInput := readUserInput(reader, "Book Price: ")

	idValue, _ := strconv.Atoi(idInput)
	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	book := NewBook(idValue, titleInput, descInput, priceValue)

	return book
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
