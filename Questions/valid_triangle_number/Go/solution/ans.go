package solution

import (
	"sort"
)

func TriangleNumber(nums []int) int {
	size := len(nums)
	count := 0
	sort.Ints(nums)

	for i := size - 1; i > 0; i-- {
		l := 0
		r := i - 1

		for l < r {
			twoSideSum := nums[l] + nums[r]
			if twoSideSum > nums[i] {
				count += (r - l)
				r--
			} else {
				l++
			}
		}
	}
	return count
}
