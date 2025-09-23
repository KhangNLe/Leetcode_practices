package solution_test

import (
	"products/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		nums []int
		ans  []int
	}{
		{
			[]int{1, 2, 4, 6},
			[]int{48, 24, 12, 8},
		},
		{
			[]int{-1, 0, 1, 2, 3},
			[]int{0, -6, 0, 0, 0},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for _, input := range inputs {
		ans := solution.ProductExceptSelf(input.nums)
		if !compareArr(ans, input.ans) {
			t.Fatalf(`
				Mismatching answer. Expected %v, got %v\n
				`, input.ans, ans)
		}
	}
}

func compareArr(a, b []int) bool {
	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
