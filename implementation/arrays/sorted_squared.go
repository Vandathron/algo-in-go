package arrays

func SortedSquaredArray(array []int) []int {
	newArray := make([]int, len(array))
	lIdx := 0
	rIdx := len(array)-1
	for idx := rIdx; idx >= 0; idx-- {
		lValue := array[lIdx]*array[lIdx]
		rValue := array[rIdx]*array[rIdx]
		if lValue >= rValue {
			newArray[idx] = lValue
			lIdx++
		} else {
			newArray[idx] = rValue
			rIdx--
		}
	}

	return newArray
}