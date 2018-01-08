package main

import (
	"fmt"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	candidate := 0
	result := PosPeaks{[]int{}, []int{}}

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			candidate = i
		}
		if array[i] < array[i-1] && candidate != 0 {
			result.Pos = append(result.Pos, candidate)
			result.Peaks = append(result.Peaks, array[candidate])
			candidate = 0
		}
	}
	return result
}

func main() {
	input := []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}
	// actual := PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}
	fmt.Println(PickPeaks(input))
	fmt.Println(PickPeaks([]int{}))
}

// inc -> newPotentialPeak
// dec -> if newPP (save it, clean it)
//
//
//
//
