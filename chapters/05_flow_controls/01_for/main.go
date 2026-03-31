package main

import "fmt"

func main() {

	total := 0
	for i := 0; i < 5; i++ {
		total += i
	}
	fmt.Printf("Sum of 0 to 4: %d\n", total)

	versions := []string{"1.18", "1.21", "1.23"}
	for idx, ver := range versions {
		fmt.Printf("Go %s is version #%d\n", ver, idx+1)
	}

	countdown := 3
	for countdown > 0 {
		fmt.Printf("%d...\n", countdown)
		countdown--
	}
	fmt.Println("Go!")
}
