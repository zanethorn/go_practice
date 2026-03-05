package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Input a string to count: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil || len(input) == 0 {
		fmt.Println("Invalid input. Please enter a non-empty string.")
		return
	}

	words := make(map[string]int)
	lowered := strings.ToLower(input)
	start := 0
	for i := 0; i < len(input); i++ {
		if lowered[i] == ' ' || i == len(input)-1 {
			snip := lowered[start:i]
			start = i + 1
			if snip != "" {
				words[snip]++
			}
		}
	}

	for k, v := range words {
		fmt.Printf("%s: %d\n", k, v)
	}
}
