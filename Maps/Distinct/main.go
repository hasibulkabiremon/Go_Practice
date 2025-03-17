package main

import "strings"

func countDistinctWords(messages []string) int {
	words := make(map[string]bool)

	for _, message := range messages {
		// Split the message into words and convert to lowercase
		messageWords := strings.Fields(strings.ToLower(message))

		// Add each word to the map
		for _, word := range messageWords {
			words[word] = true
		}
	}

	return len(words)
}
