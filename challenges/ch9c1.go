package main

import "strings"

func countDistinctWords(messages []string) int {

	distinctWords := make(map[string]int)

	for _, message := range messages {
		message = strings.ToLower(message)
		words := strings.Fields(message)
		for _, word := range words {
			distinctWords[word]++
		}
	}

	return len(distinctWords)
}
