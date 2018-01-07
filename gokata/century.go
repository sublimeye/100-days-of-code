package main

import (
	"fmt"
)

func century(year int) int {
	cent := year / 100
	if year == cent*100 {
		return cent
	}

	return cent + 1
}

// codewars better / shorter / simplier* solution
func centurySimpleMath(year int) int {
	return (year + 99) / 100
}

// codewars better / shorter / simplier* solution
func century2(year int) int {
	return (year-1)/100 + 1
}

func main() {
	fmt.Println(century(1900)) // 19
	fmt.Println(century(1990)) // 20
	fmt.Println(century(2001)) // 21

	fmt.Println(centurySimpleMath(1900)) // 19
	fmt.Println(centurySimpleMath(1990)) // 20
	fmt.Println(centurySimpleMath(2001)) // 21

	fmt.Println(century2(1900)) // 19
	fmt.Println(century2(1990)) // 20
	fmt.Println(century2(2001)) // 21
}
