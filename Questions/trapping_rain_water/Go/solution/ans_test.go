package solution_test

import (
	"testing"
	"trapping/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		height []int
		ans    int
	}{
		{
			[]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			9,
		},
		{
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
	}

	for _, input := range inputs {
		res := solution.Trap(input.height)
		if res != input.ans {
			t.Fatalf(`
				Mismatching input. Expected %d, got %d
				`, input.ans, res)
		}
	}
}
