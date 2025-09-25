package solution_test

import (
	"reflect"
	"sum/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		num    []int
		target int
		res    []int
	}{
		{
			[]int{1, 2, 3, 4},
			3,
			[]int{1, 2},
		},
		{
			[]int{-4, -3, -2, -1, 0, 2, 6},
			4,
			[]int{3, 7},
		},
		{
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
	}

	for _, input := range inputs {
		ans := solution.TwoSum(input.num, input.target)
		if !reflect.DeepEqual(ans, input.res) {
			t.Fatalf(`
				Mismatching result. Expect %v, got %v
				`, input.res, ans)
		}
	}
}
