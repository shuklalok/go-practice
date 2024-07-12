package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a string")
		return
	}

	str := os.Args[1]

	fmt.Printf("Total characters in the Original string: %v\n", len(str))
	fmt.Printf("Number of vowels: %v\n", countVowels(str))
	fmt.Printf("Number of consonants: %v\n", countConsonants(str))
	fmt.Printf("Number of numerics: %v\n", countNumerics(str))
	fmt.Printf("Number of spaces: %v\n", countSpaces(str))
	fmt.Printf("Number of special characters: %v\n", countSpecialCharacters(str))
}

func countVowels(str string) int {
	return strings.Count(str, "a") + strings.Count(str, "e") + strings.Count(str, "i") + strings.Count(str, "o") + strings.Count(str, "u")
}

func countConsonants(str string) int {
	return len(str) - countVowels(str) - countNumerics(str) - countSpaces(str) - countSpecialCharacters(str)
}

func countNumerics(str string) int {
	count := 0
	for _, c := range str {
		if c >= '0' && c <= '9' {
			count++
		}
	}
	return count
}

func countSpaces(str string) int {
	return strings.Count(str, " ")
}

func countSpecialCharacters(str string) int {
	specialCharacters := map[rune]bool{
		'!': true, '@': true, '#': true, '$': true, '%': true,
		'^': true, '&': true, '*': true, '(': true, ')': true,
	}
	count := 0
	for _, c := range str {
		if specialCharacters[c] {
			count++
		}
	}
	return count
}
