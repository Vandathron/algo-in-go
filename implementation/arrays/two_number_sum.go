package arrays

func Sol1(array []int, target int)[]int{
	subs := make(map[int]int)
	for _, val := range array  {
		s := target-val
		if _,ok := subs[s]; ok {
			return []int {s,val}
		}
		subs[val] = 0
	}
	return []int {}
}
