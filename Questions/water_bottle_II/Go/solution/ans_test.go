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
			13,
			6,
			15,
		},
		{
			10,
			3,
			13,
		},
		{
			10,
			1,
			14,
		},
	}

	for _, input := range inputs {
		ans := solution.MaxBottlesDrunk(input.bottles, input.exchange)
		if ans != input.res {
			t.Fatalf(`
				Mismatching answer. Expected %d, got %d
				`, input.res, ans)
		}
	}
}
