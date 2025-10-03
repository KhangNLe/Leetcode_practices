package solution_test

import (
	"container/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		heights []int
		ans     int
	}{
		{
			[]int{1, 7, 2, 5, 4, 7, 3, 6},
			36,
		},
		{
			[]int{2, 2, 2},
			4,
		},
		{
			[]int{1, 2, 1},
			2,
		},
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 1},
			1,
		},
	}

	for _, input := range inputs {
		res := solution.MaxArea(input.heights)
		if res != input.ans {
			t.Fatalf(`
				Mismatching answer. Expected %d, got %d
				`, input.ans, res)
		}
	}
}
