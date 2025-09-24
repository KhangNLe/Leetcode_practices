package solution_test

import (
	"fraction/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		numberator  int
		denominator int
		res         string
	}{
		{
			1,
			2,
			"0.5",
		},
		{
			2,
			1,
			"2",
		},
		{
			4,
			333,
			"0.(012)",
		}, {
			1,
			3,
			"0.(3)",
		}, {
			113,
			1000,
			"0.113",
		},
		{
			111111111234,
			1000000000000,
			"0.111111111234",
		},
		{
			0,
			-1,
			"0",
		},
		{
			-3,
			-4,
			"0.75",
		},
		{
			1,
			-4,
			"-0.25",
		},
	}

	for _, input := range inputs {
		ans := solution.FractionToDecimal(input.numberator, input.denominator)
		if ans != input.res {
			t.Fatalf(`
				Mismatching result. Expect %s, got %s
				`, input.res, ans)
		}
	}
}
