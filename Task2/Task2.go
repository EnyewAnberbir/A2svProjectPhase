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
func IsPalindrome(input string) bool {

	input = strings.ToLower(input)  	
	re := regexp.MustCompile(`[^\w]`)
	input = re.ReplaceAllString(input, "")

	len_ := len(input)
	for i := 0; i < len_/2; i++ {
		if input[i] != input[len_-1-i] {
			return false
		}
	}
	return true
}
func main() {
	text := "Artificial intelligence and machine learning are transforming the future of technology, as artificial intelligence continues to advance"
	freq := freqCount(text)
	fmt.Println(freq)
	text1 := "No lemon, no melon"
	isPalindrome := IsPalindrome(text1)
	fmt.Printf("Is the text \"%s\" a palindrome? %v\n", text1, isPalindrome)
}
