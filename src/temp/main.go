package main

import (
	"fmt"
	"strconv"
)

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func CtoF(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	fmt.Print("Input Temp and scale (e.g. 100F or 0C): ")
	var input string
	fmt.Scanln(&input)

	length := len(input)
	temp, err := strconv.ParseFloat(input[:length-1], 64)
	if err != nil {
		fmt.Println("Invalid temperature value. Please enter a valid number followed by 'F' or 'C'.")
		return
	}
	scale := input[length-1:]

	switch scale {
	case "F", "f":
		fmt.Printf("%g°F is %g°C\n", temp, FtoC(temp))
	case "C", "c":
		fmt.Printf("%g°C is %g°F\n", temp, CtoF(temp))
	default:
		fmt.Println("Invalid scale. Please use 'F' for Fahrenheit or 'C' for Celsius.")
	}

}
