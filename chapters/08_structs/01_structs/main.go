package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages
}

func main() {

	book := {
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Pages:  380,
	}

	fmt.Printf("%s by %s (%d pages)\n", book.Title, book.Author, book.Pages)

	book. = 400

	fmt.Printf("Updated pages: %d\n", book.Pages)
}
