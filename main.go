package main

import (
	"fmt"
)

func Calculate(x int) (result int) {
	result = x + x

	return result
}

func main() {
	fmt.Printf("Hello World, %d", Calculate(1))
}
