package arrays

func MoveElementToEnd(array []int, toMove int) []int {
	l := 0
	r := len(array)-1
	for l < r {
		if array[l] == toMove {
			if array[r] == toMove {
				r-=1
			} else {
				temp := array[r]
				array[r] = array[l]
				array[l] = temp
				l+=1
			}
		} else {
			l+=1
		}
	}
	return array
}