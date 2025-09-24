package solution

func LongestConsecutive(nums []int) int {
	numFreq := make(map[int]struct{})

	for _, num := range nums {
		numFreq[num] = struct{}{}
	}

	longest := 0

	for num, _ := range numFreq {
		if _, ok := numFreq[num-1]; !ok {
			curr := 1
			for {
				if _, ok = numFreq[num+curr]; ok {
					curr++
				} else {
					break
				}
			}
			longest = max(longest, curr)
		}
	}
	return longest
}
