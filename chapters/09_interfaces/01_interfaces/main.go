package main

import "fmt"

type Animal interface {
	Describe() string
}

type Dog struct {
	Name string
}

func (d Dog) Describe() string {
	return fmt.Sprintf("Dog: %s", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Describe() string {
	return fmt.Sprintf("Cat: %s", c.Name)
}

func printAnimal(a Animal) {
	fmt.Println(a.Describe())
}

func main() {

	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}

	printAnimal(dog)
	printAnimal(cat)
}
