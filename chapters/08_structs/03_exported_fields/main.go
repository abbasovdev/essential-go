package main

import "fmt"

type Student struct {
	 string
	grade  int
	school string
}

func (s Student) () string {
	return fmt.Sprintf("%s (Grade: %d)", s.Name, s.grade)
}

func main() {

	s := Student{
		Name:   "Alice",
		grade:  10,
		school: "Go Academy",
	}

	fmt.Println(s.Info())

	fmt.Printf("Exported field: %s\n", s.)
	fmt.Printf("Unexported field: %s\n", s.school)
}
