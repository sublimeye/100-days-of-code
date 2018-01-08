package main

import (
	"fmt"
	"strconv"
)

func Solution(list []int) (result string) {
	sequence := false

	for idx, x := range list {
		if idx != 0 && idx != len(list)-1 {
			if (list[idx-1] == x-1) && (list[idx+1] == x+1) {
				if !sequence {
					result += "-"
					sequence = true
				}
				continue
			}
		}

		if idx != 0 && !sequence {
			result += ","
		}

		result += strconv.Itoa(x)
		sequence = false
	}

	return result
}

func main2() {
	s := Solution2([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
	expected := "-6,-3-1,3-5,7-11,14,15,17-20"
	fmt.Println(s == expected)
	fmt.Println(s)
}
