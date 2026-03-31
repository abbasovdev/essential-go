package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.("something went wrong")

	fmt.Println(err)

	name := "config.yaml"
	err = fmt.Errorf("file not found: %s", )

	fmt.Println(err)
}
