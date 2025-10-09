package solution_test

import (
	"potion/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		skill []int
		mana  []int
		ans   int64
	}{
		{
			[]int{1, 5, 2, 4},
			[]int{5, 1, 4, 2},
			110,
		},
		{
			[]int{1, 1, 1},
			[]int{1, 1, 1},
			5,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2},
			21,
		},
	}

	for _, input := range inputs {
		ans := solution.MinTime(input.skill, input.mana)
		if ans != input.ans {
			t.Fatalf(`
				Mismatching answer. Expected %d, got %d
				`, input.ans, ans)
		}
	}
}
