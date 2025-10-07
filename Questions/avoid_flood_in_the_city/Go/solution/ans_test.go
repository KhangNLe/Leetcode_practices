package solution_test

import (
	"flood/solution"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		rains []int
		ans   []int
	}{
		{
			[]int{1, 0, 2, 0, 3, 0, 2, 0, 0, 0, 1, 2, 3},
			[]int{-1, 1, -1, 2, -1, 3, -1, 2, 1, 1, -1, -1, -1},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{-1, -1, -1, -1},
		},
		{
			[]int{1, 2, 0, 0, 2, 1},
			[]int{-1, -1, 2, 1, -1, -1},
		},
		{
			[]int{1, 2, 0, 1, 2},
			[]int{},
		},
		{
			[]int{0, 1, 1},
			[]int{},
		},
		{
			[]int{0, 0, 0, 1, 2, 0, 2},
			[]int{1, 1, 1, -1, -1, 2, -1},
		},
		{
			[]int{1, 0, 2, 3, 0, 1, 2},
			[]int{-1, 1, -1, -1, 2, -1, -1},
		},
	}

	for _, input := range inputs {
		ans := solution.AvoidFlood(input.rains)
		if !reflect.DeepEqual(ans, input.ans) {
			t.Fatalf(`
				Mismatching answer. Expected %v, got %v
				`, input.ans, ans)
		}
	}
}
