package solution

func Trap(height []int) int {
	secondHighest := 0
	l := 0
	r := len(height) - 1
	total := 0
	for l < r {
		smallerLine := min(height[l], height[r])
		if smallerLine < secondHighest {
			total += (secondHighest - smallerLine)
		} else {
			secondHighest = smallerLine
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return total
}
