package solution

func MaxArea(heights []int) int {
	maxScore := 0
	l := 0
	r := len(heights) - 1

	for l < r {
		smallest := min(heights[l], heights[r])

		maxScore = max(maxScore, smallest*(r-l))

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return maxScore
}
