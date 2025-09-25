package solution_test

import (
	"testing"
	"triangle/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		triangle [][]int
		ans      int
	}{
		{
			[][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			11,
		},
		{
			[][]int{{-10}},
			-10,
		},
		{
			[][]int{
				{2},
				{2, 5},
				{7, 9, 4},
				{6, 9, 2, 6},
				{3, 7, 9, 2, 7},
			},
			15,
		},
	}

	for _, input := range inputs {
		ans := solution.MinimumTotalBottomUp(input.triangle)
		if ans != input.ans {
			t.Fatalf(`
				Mismatching input. Expect %d got %d
				`, input.ans, ans)
		}
	}
}
