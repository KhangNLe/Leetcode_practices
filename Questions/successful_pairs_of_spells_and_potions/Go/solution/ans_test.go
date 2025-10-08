package solution_test

import (
	"reflect"
	"sort"
	"spell/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		spells  []int
		potions []int
		succ    int64
		ans     []int
	}{
		{
			[]int{15, 8, 19},
			[]int{38, 36, 23},
			328,
			[]int{3, 0, 3},
		},
		{
			[]int{3, 1, 2, 3, 5, 236, 1},
			[]int{2, 2, 2, 2, 2, 54, 7, 2, 12, 12, 13, 12, 12, 12, 45},
			16,
			[]int{9, 2, 8, 9, 9, 15, 2},
		},
		{
			[]int{5, 1, 3},
			[]int{1, 2, 3, 4, 5},
			7,
			[]int{4, 0, 3},
		},
		{
			[]int{3, 1, 2},
			[]int{8, 5, 8},
			16,
			[]int{2, 0, 2},
		},
	}

	for _, input := range inputs {
		ans := solution.SuccessfulPairs(input.spells, input.potions, input.succ)
		sort.Ints(ans)
		sort.Ints(input.ans)
		if !reflect.DeepEqual(ans, input.ans) {
			t.Fatalf(`
				Mismatching answer. Expected %v, got %v
				`, input.ans, ans)
		}
	}
}
