package main

import (
	"fmt"
	"regexp"
	"strings"
)

// IsPalindrome checks whether a given string is a palindrome
func IsPalindrome(input string) bool {
	// Define a regex pattern to match only alphabetic characters
	regex := regexp.MustCompile(`[a-zA-Z0-9]+`)

	// Find all matches and concatenate them into a single string
	processedString := strings.Join(regex.FindAllString(strings.ToLower(input), -1), "")

	// Initialize pointers for the start and end of the string
	start, end := 0, len(processedString)-1

	// Compare characters from the start and end moving towards the center
	for start < end {
		if processedString[start] != processedString[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	// Take user input
	fmt.Println("Enter a string to check if it is a palindrome:")
	var input string
	fmt.Scanln(&input)

	// Check if the input string is a palindrome
	if IsPalindrome(input) {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
