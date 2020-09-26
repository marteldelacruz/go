package main

import (
	"fmt"
)

func main() {
	// store fahrenheit value as float
	var fahrenheit float64

	// Banner text
	fmt.Println("----------Fahrenheit to Celsius-----------")

	// Ask user for the value
	fmt.Print("\nFahrenheit: ")
	fmt.Scanf("%g", &fahrenheit)

	// Print result
	fmt.Printf("\nCelsius: %g", (fahrenheit-32)*5/9)
}
