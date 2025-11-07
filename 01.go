package main

import "strings"

// Program to find longest word in a given sentence
func findLongestWord(sentence string) string {
	if sentence == "" {
		return ""
	}
	sentenceArr := strings.Split(sentence, " ")
	longestSentenceLength := 0
	longestWord := ""
	for _, word := range sentenceArr {
		if len(word) >= longestSentenceLength {
			longestWord = word
			longestSentenceLength = len(word)
		}
	}
	return longestWord
}
