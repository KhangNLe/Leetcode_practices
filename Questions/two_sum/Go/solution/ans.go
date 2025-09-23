package solution

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		numMap[num] = i
	}

	for i, num := range nums {
		remain := target - num
		if idx, ok := numMap[remain]; ok {
			if idx != i {
				return []int{idx, i}
			}
		}
	}
	return nil
}
