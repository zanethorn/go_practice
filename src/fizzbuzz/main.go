package main

import "fmt"

func FizzBuzz(n int) []string {
	result := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result[i] = "FizzBuzz"
		} else if i%5 == 0 {
			result[i] = "Buzz"
		} else if i%3 == 0 {
			result[i] = "Fizz"
		} else {
			result[i] = ""
		}
	}
	return result[1:]
}

func main() {
	fmt.Println(FizzBuzz(45))
}
