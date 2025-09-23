package solution_test

import (
	"frequence/solution"
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{
			[]int{1, 2, 2, 3, 3, 3},
			2,
			[]int{2, 3},
		},
		{
			[]int{7, 7},
			1,
			[]int{7},
		},
		{
			[]int{1, 2, 3, 1, 1, 1, 1, 2, 6, 3, 3, 4, 5, 2, 2, 5, 6},
			3,
			[]int{1, 2, 3},
		}, {
			[]int{-1, -1},
			1,
			[]int{-1},
		}, {
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
	}

	for _, input := range inputs {
		ans := solution.TopKFrequent(input.nums, input.k)
		sort.Ints(ans)
		sort.Ints(input.expect)

		if !compareArr(ans, input.expect) {
			t.Fatalf(`
				Mismatching array. Expected %v, got %v\n
				`, input.expect, ans)
		}
	}
}

func compareArr(a, b []int) bool {
	n := len(a)
	m := len(b)

	if n != m {
		return false
	}

	for i := range n {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
