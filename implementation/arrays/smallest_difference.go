package arrays

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)

	smallestDiff := math.MaxInt32
	pair := []int{}
	l := 0
	r := 0

	for l < len(array1) && r < len(array2) {
		lVal := array1[l]
		rVal := array2[r]

		diff := lVal-rVal
		if lVal < rVal {
			l+=1
		} else if lVal > rVal {
			r+=1
		} else {
			return []int {lVal, rVal}
		}
		if diff < 0 {
			diff *=-1
		}

		if diff < smallestDiff {
			smallestDiff = diff
			pair = []int{lVal, rVal}
		}
	}
	return pair
}