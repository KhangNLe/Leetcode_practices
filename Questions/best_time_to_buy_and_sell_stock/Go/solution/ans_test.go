package solution_test

import (
	"stock/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		prices []int
		ans    int
	}{
		{
			[]int{10, 1, 5, 6, 7, 1},
			6,
		},
		{
			[]int{10, 8, 7, 5, 2},
			0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
	}

	for _, input := range inputs {
		ans := solution.MaxProfit(input.prices)
		if ans != input.ans {
			t.Fatalf(`
				Mismatching answer. Expected %d, got %d
				`, input.ans, ans)
		}
	}
}
