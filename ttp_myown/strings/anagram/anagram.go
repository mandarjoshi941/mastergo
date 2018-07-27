package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	str1 := "aaaataArmy"
	str2 := "Maryasaaaa"
	start := time.Now()
	if isAnagram(str1, str2) {
		fmt.Println("these are anagrams")
	} else {
		fmt.Println("these are not anagrams")
	}
	elapsed := time.Since(start)
	fmt.Println("Time Taken:", elapsed)

	start = time.Now()
	isAnagramOldWay(str1, str2)
	elapsed = time.Since(start)
	fmt.Println("Time Taken:", elapsed)
}

func isAnagram(str1 string, str2 string) bool {
	charMap1 := getCharCount(strings.ToLower(str1))
	charMap2 := getCharCount(strings.ToLower(str2))
	return compareCharMaps(charMap1, charMap2)
}

func compareCharMaps(map1 map[rune]int, map2 map[rune]int) bool {
	for k, v := range map1 {
		if map2[k] == 0 || map2[k] != v {
			return false
		}
	}
	return true
}
func getCharCount(str string) map[rune]int {
	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
	}
	return charMap
}

func isAnagramOldWay(str1 string, str2 string) {
	lowerStr1 := strings.ToLower(str1)
	lowerStr2 := strings.ToLower(str2)
	runeArrayStr1 := []rune(lowerStr1)
	runeArrayStr2 := []rune(lowerStr2)

	for idx, char := range lowerStr1 {
		index := strings.IndexRune(lowerStr2, char)
		if index != -1 {
			runeArrayStr2 = removeFromArray(runeArrayStr2, index)
			lowerStr2 = string(runeArrayStr2)
			runeArrayStr1 = removeFromArray(runeArrayStr1, idx)
			//lowerStr1 = string(runeArrayStr1)
		}
	}

	fmt.Println("Anagarms :", len(runeArrayStr2) == 0)
}

func removeFromArray(s []rune, i int) []rune {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
