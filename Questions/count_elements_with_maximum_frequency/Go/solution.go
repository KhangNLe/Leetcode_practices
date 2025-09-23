package main

func maxFrequencyElements(nums []int) int {
	frequences := make(map[int]int)
	for _, num := range nums {
		_, ok := frequences[num]
		if !ok {
			frequences[num] = 1
		} else {
			frequences[num]++
		}
	}

	var highest_freq int
	var total int

	for _, val := range frequences {
		if highest_freq < val {
			total = val
			highest_freq = val
		} else if highest_freq == val {
			total += val
		}
	}
	return total
}
