package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {

	book := Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  380,
	}

	fmt.Printf("%s by %s (%d pages)\n", book.Title, book.Author, book.Pages)

	book.Pages = 400

	fmt.Printf("Updated pages: %d\n", book.Pages)
}
