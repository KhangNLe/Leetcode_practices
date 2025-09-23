package solution

func HasDuplicate(nums []int) bool {
	unique := make(map[int]int)
	for _, num := range nums {
		_, ok := unique[num]
		if !ok {
			unique[num] = 1
		} else {
			return true
		}
	}
	return false
}
