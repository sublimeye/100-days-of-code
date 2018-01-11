package main

import (
	"fmt"
)

func Exercise43() {

}

func Exercise44() {
	fmt.Println("Exercise44")
	slice := []int{2, 3, 4, 5, 6, 7}
	// 5, 6, 7, 2, 3, 4
	// 6, 7, 2, 3, 4, 5
	newslice := rotate(slice, 0)
	fmt.Println(newslice, slice)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, shift int) []int {
	length := len(s)
	newslice := []int{}

	if shift > 0 {
		if shift > length {
			shift = length
		}

		newslice = append(s[shift:], s[:shift-1]...)
	} else {
		if shift < -length {
			shift = -length
		}
		newslice = append(s[length+shift:], s[:length+shift]...)
	}

	return newslice
}

/**
* Exercise 4.5
*
 */
func Exercise45() {
	// test := []string{"a", "b", "b", "b", "c", "d", "d", "d", "d", "d", "d", "d", "d", "e", "d", "d", "d", "d", "d", "d"}
	test2 := []string{}
	// a, b, c, d
	newtest := removeAdjacent(test2)
	fmt.Println("Exercise45", newtest)
}

// Find adjacent duplicates in a string slice
func removeAdjacent(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}
	// go through strings
	// if prev == current => remove prev || current

	out := strings[0:1] // zero-length slice of original

	// 0,1,2,3,4,5,6
	for i, j := 1, 1; i < len(strings); i++ {
		if strings[i] != out[j-1] {
			out = append(out, strings[i])
			j++
		}
	}

	return out
}

func main() {
	// Exercise43()
	// Exercise44()
	Exercise45()
}
