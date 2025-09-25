package solution

func MinimumTotalBottomUp(triangle [][]int) int {
	height := len(triangle)
	base := len(triangle[height-1])

	table := make([]int, base)
	copy(table, triangle[height-1])

	for i := height - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			table[j] = triangle[i][j] + min(table[j], table[j+1])
		}
	}

	return table[0]
}
