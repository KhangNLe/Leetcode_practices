package solution

func TopKFrequent(nums []int, k int) []int {
	numFreq := make(map[int]int)

	for _, num := range nums {
		numFreq[num]++
	}

	freqLayout := make([][]int, len(nums)+1)

	for num, freq := range numFreq {
		freqLayout[freq] = append(freqLayout[freq], num)
	}

	ans := []int{}

	for i := len(freqLayout) - 1; i >= 0; i-- {
		if k == 0 {
			return ans
		}
		size := len(freqLayout[i])
		for size > 0 && k > 0 {
			j := size - 1
			ans = append(ans, freqLayout[i][j])
			size--
			k--
		}
	}
	return nil
}
