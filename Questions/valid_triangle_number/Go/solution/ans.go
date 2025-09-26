package solution

import (
	"sort"
)

func TriangleNumber(nums []int) int {
	size := len(nums)
	if size <= 2 {
		return 0
	}

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
