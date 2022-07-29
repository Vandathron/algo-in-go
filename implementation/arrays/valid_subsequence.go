package arrays


func sol1(array []int, sequence []int) bool {
	idx := 0
	for _, a := range array {

		if a == sequence[idx]{
			idx++
		}

		if idx == len(sequence) {
			break
		}
	}
	return idx == len(sequence)
}