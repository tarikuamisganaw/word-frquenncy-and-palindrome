package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// WordFrequencyCount takes a string as input and returns a map with the frequency of each word
func WordFrequencyCount(input string) map[string]int {
	// Define a regex pattern to match words
	regex := regexp.MustCompile(`[a-zA-Z]+`)

	// Find all words in the input string and convert to lowercase
	words := regex.FindAllString(strings.ToLower(input), -1)

	// Initialize a map to hold word frequencies
	wordFreq := make(map[string]int)

	// Iterate over the words and count frequencies
	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

// IsPalindrome checks whether a given string is a palindrome
func IsPalindrome(input string) bool {
	// Function to remove non-alphanumeric characters and convert to lowercase
	cleanString := func(str string) string {
		var sb strings.Builder
		for _, r := range str {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				sb.WriteRune(unicode.ToLower(r))
			}
		}
		return sb.String()
	}

	cleanedInput := cleanString(input)
	length := len(cleanedInput)
	for i := 0; i < length/2; i++ {
		if cleanedInput[i] != cleanedInput[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Word Frequency Count
	fmt.Println("Enter a string to analyze word frequency:")
	input, _ := reader.ReadString('\n')
	wordFreq := WordFrequencyCount(strings.TrimSpace(input))

	fmt.Println("Word Frequencies:")
	for word, count := range wordFreq {
		fmt.Printf("%s: %d\n", word, count)
	}

	// Palindrome Check
	fmt.Println("\nEnter a string to check if it's a palindrome:")
	input, _ = reader.ReadString('\n')
	isPalindrome := IsPalindrome(strings.TrimSpace(input))

	if isPalindrome {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
