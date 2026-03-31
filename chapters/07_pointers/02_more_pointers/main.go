package main

import "fmt"

func addBonus(score int, bonus int) {
	score = score + bonus
}

func main() {

	score := 50
	bonus := 25

	addBonus(score, bonus)

	fmt.Printf("Score after bonus: %d\n", score)
}
