package solution

import (
	"math"
)

func MinimumTotalTopDown(triangle [][]int) int {
	height := len(triangle)
	base := len(triangle[height-1])

	table := make([]int, base)

	for i := range height {
		data := make([]int, base)
		if i == 0 {
			data[0] = triangle[0][0]
		} else {
			for j := range i + 1 {
				switch j {
				case 0:
					data[j] = triangle[i][j] + table[j]
				case i:
					data[j] = triangle[i][j] + table[j-1]
				default:
					data[j] = triangle[i][j] + min(table[j], table[j-1])
				}
			}
		}
		table = data
	}

	smallest := math.MaxInt

	for _, num := range table {
		smallest = min(smallest, num)
	}
	return smallest
}
