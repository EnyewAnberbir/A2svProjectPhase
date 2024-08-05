package main

import (
	"fmt"
	"regexp"
	"strings"
)

func freqCount(input string) map[string]int {
	input = strings.ToLower(input)
	re := regexp.MustCompile(`[^\w\s]`)
	input = re.ReplaceAllString(input, "")
	words := strings.Fields(input)

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

func main() {
	text := "Hello, A2SVian! Hello, enyew. This is a task 2, task 2 completed."
	freq := freqCount(text)
	fmt.Println(freq)
}
