package main

import "fmt"

func main() {

	var language string = "Go"

	var version float64 = 1.23

	var year int = 2009

	var isOpenSource bool = true

	versionAsInt := int(version)

	fmt.Printf("%s version %.2f was released in %d. Open source: %t. Major version: %d\n", language, version, year, isOpenSource, versionAsInt)
}
