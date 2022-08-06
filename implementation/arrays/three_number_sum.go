package arrays

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	sums := [][]int{}

	for i := 0; i < len(array)-2; i++ {
		cur := array[i]
		l := i+1
		r := len(array)-1

		for l < r{
			sum := cur + array[l] + array[r]
			if sum == target {
				sums = append(sums, []int{cur, array[l], array[r]})
				l+=1
				r-=1
			} else if sum < target {
				l+=1
			} else {
				r -=1
			}
		}
	}
	return sums
}