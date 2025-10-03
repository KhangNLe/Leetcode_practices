package solution_test

import (
	"testing"
	"trap/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		heightMap [][]int
		ans       int
	}{
		{
			[][]int{
				{12, 13, 1, 12},
				{13, 4, 13, 12},
				{13, 8, 10, 12},
				{12, 13, 12, 12},
				{13, 13, 13, 13},
			},
			14,
		},
		{
			[][]int{
				{1, 4, 3, 1, 3, 2},
				{3, 2, 1, 3, 2, 4},
				{2, 3, 3, 2, 3, 1},
			},
			4,
		},
		{
			[][]int{
				{3, 3, 3, 3, 3},
				{3, 2, 2, 2, 3},
				{3, 2, 1, 2, 3},
				{3, 2, 2, 2, 3},
				{3, 3, 3, 3, 3},
			},
			10,
		},
	}

	for _, input := range inputs {
		res := solution.TrapRainWater(input.heightMap)
		if res != input.ans {
			t.Fatalf(`
				Mismatching answer. Expected %d, got %d
				`, input.ans, res)
		}
	}
}
