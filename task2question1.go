package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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

func main() {
	// Take user input
	fmt.Println("Enter a string to analyze word frequency:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Get word frequency count
	wordFreq := WordFrequencyCount(strings.TrimSpace(input))

	// Print the word frequencies
	fmt.Println("Word Frequencies:")
	for word, count := range wordFreq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
