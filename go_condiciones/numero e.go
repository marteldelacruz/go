package main

import "fmt"

// Returns the facturial of a positive number
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	// Store result
	var e float64
	// Iterations
	var totalIterations int = 40

	for i := 0; i < totalIterations; i++ {
		e += 1 / float64(factorial(i))
	}

	fmt.Print(e)
}
