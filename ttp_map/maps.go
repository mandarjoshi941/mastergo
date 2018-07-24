package main

import (
	"fmt"
	"strings"
)

func main() {
	mapVar := getWordCount("my name is mandar, And my full name is Mandar anand joshi")
	printMap(mapVar)
}

func getWordCount(str string) map[string]int {
	wordCountmap := make(map[string]int)
	stringSlice := strings.Split(str, " ")
	for _, word := range stringSlice {
		trimmedWord := strings.Trim(word, " '.,:;?!()-[]<>{}\t\n\"")
		lowerWord := strings.ToLower(trimmedWord)
		wordCountmap[lowerWord]++
	}
	return wordCountmap
}

func printMap(mapArg map[string]int) {
	for k, v := range mapArg {
		if v > 1 {
			fmt.Printf("Key:%10s Value: %2d\n", k, v)
		}
	}
}
