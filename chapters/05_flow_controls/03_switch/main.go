package main

import "fmt"

func main() {

	language := "Go"
	 language {
	case "Go":
		fmt.Println("Go: compiled and fast")
	case "Python":
		fmt.Println("Python: interpreted and flexible")
	:
		fmt.Println("Unknown language")
	}

	year := 2009
	switch {
	 year < 2000:
		fmt.Println("Created before 2000")
	case year >= 2000:
		fmt.Println("Created in 2000 or later")
	}

	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	}
}
