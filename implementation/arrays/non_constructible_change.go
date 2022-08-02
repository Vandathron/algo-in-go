package arrays

import "sort"

func NonConstructibleChange(coins []int) int {

	currChange := 0
	sort.Ints(coins)

	for _, coin := range coins {
		if coin > currChange + 1 {
			return currChange +1
		}
		currChange += coin
	}
	return currChange + 1
}