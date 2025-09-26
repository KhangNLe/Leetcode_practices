package solution_test

import (
	"testing"
	"triangle/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		nums []int
		res  int
	}{
		{
			[]int{48, 66, 61, 46, 94, 75},
			19,
		},
		{
			[]int{2, 2, 3, 4},
			3,
		},
		{
			[]int{4, 2, 3, 4},
			4,
		},
		{
			[]int{2, 3},
			0,
		},
		{
			[]int{3, 4, 5},
			1,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			7,
		},
	}

	for _, input := range inputs {
		ans := solution.TriangleNumber(input.nums)
		if ans != input.res {
			t.Fatalf(`
				Mismatching result. Expect %d, got %d
				`, input.res, ans)
		}
	}
}
