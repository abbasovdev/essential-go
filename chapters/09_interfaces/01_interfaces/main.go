package main

import "fmt"

type  interface {
	Describe() string
}

type Dog struct {
	Name string
}

func (d Dog) () string {
	return fmt.Sprintf("Dog: %s", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Describe()  {
	return fmt.Sprintf("Cat: %s", c.Name)
}

func printAnimal(a ) {
	fmt.Println(a.Describe())
}

func main() {

	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}

	printAnimal()
	printAnimal()
}
