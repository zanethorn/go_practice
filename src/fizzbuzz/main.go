package main

import "fmt"

func FizzBuzz(n int) []string {
	result := make([]string, n+1)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i] = "FizzBuzz"
		case i%5 == 0:
			result[i] = "Buzz"
		case i%3 == 0:
			result[i] = "Fizz"
		default:
			result[i] = ""
		}
	}
	return result[1:]
}

func main() {
	fmt.Println(FizzBuzz(45))
}
