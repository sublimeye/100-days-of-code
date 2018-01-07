package main

import (
	"fmt"
	"strings"
)

// implemented by the noob Roman, first attempt
func duplicateCount(s1 string) int {
	charCounter := make(map[string]int)
	var max int

	for _, char := range s1 {
		charCounter[strings.ToUpper(string(char))]++
	}

	for _, value := range charCounter {
		if value > 1 {
			max++
		}
	}

	return max
}

// implemented by some guy on codewars
// 1. strings.ToLower accept rune
// 2. convert s string to lowercase when ranging over elements
// 3. use "pre" assignment in if statement
// 4. use a custom variable for the return value
func bestPracticeDuplicateCount(s string) (counter int) {
	hash := make(map[rune]int)
	for _, r := range strings.ToLower(s) {
		if hash[r]++; hash[r] == 2 {
			counter++
		}
	}

	return counter
}

func RunCountingDuplicates() {
	s1 := "aaaaaddddddddddkkkkkkk123123123123123123"
	s1Count := bestPracticeDuplicateCount(s1)
	fmt.Printf("Counted %d duplicates in \"%v\"", s1Count, s1)
}
