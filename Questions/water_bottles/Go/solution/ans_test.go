package solution_test

import (
	"testing"
	"water/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		bottles  int
		exchange int
		res      int
	}{
		{
			9,
			3,
			13,
		},
		{
			15,
			4,
			19,
		},
	}

	for _, input := range inputs {
		ans := solution.NumWaterBottles(input.bottles, input.exchange)
		if ans != input.res {
			t.Fatalf(`
				Mismatching result. Expect %d, got %d
				`, input.res, ans)
		}
	}
}
