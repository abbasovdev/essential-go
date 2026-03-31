package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.New("something went wrong")

	fmt.Println(err)

	name := "config.yaml"
	err = fmt.Errorf("file not found: %s", name)

	fmt.Println(err)
}
