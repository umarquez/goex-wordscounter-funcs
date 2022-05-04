package main

import (
	"fmt"
	"regexp"
	"strings"
	"words_counter/services/random_text_api"
	"words_counter/services/static_text"
)

func extractWordsDeprecated(text string) []string {
	words := strings.Split(text, " ")
	fmt.Printf("total words: %v\n\n", len(words))
	return words
}

func extractWords(text string) []string {
	words := regexp.MustCompile(`[a-zA-Z/-]+`).FindAllString(text, -1)
	fmt.Printf("total words: %v\n\n", len(words))
	return words
}

func countWordsOccurrences(words ...string) (wordsCounters map[string]int) {
	wordsCounters = make(map[string]int)
	for _, word := range words {
		wordsCounters[strings.ToLower(word)]++
	}

	return
}

func printMapContent(content map[string]int) {
	for key, val := range content {
		fmt.Printf("%v - %v\n", key, val)
	}
}

func main() {
	txt := static_text.GetText
	wordList := extractWords(txt())
	wordsCounters := countWordsOccurrences(wordList...)
	printMapContent(wordsCounters)

	fmt.Println("\n------------------\n")
	// changing source
	txt = random_text_api.GetText
	// single line call
	printMapContent(countWordsOccurrences(extractWords(txt())...))
}
