package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {

	rect := Rectangle{Width: 5.0, Height: 3.0}

	fmt.Printf("Area: %.1f\n", rect.Area())

	rect.Scale(2.0)

	fmt.Printf("Scaled: %.1f x %.1f\n", rect.Width, rect.Height)
}
