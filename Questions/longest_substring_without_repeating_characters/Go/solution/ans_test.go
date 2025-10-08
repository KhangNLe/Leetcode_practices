package solution_test

import (
	"substring/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		s   string
		ans int
	}{
		{
			"abcbefga",
			6,
		},
		{
			"dvdfg",
			4,
		},
		{
			"ab",
			2,
		},
		{
			"zxyzxyzxy",
			3,
		},
		{
			"xxxxx",
			1,
		},
		{
			"abcbcdefgd",
			6,
		},
	}

	for _, input := range inputs {
		ans := solution.LengthOfLongestSubstring(input.s)
		if ans != input.ans {
			t.Fatalf(`
				Mismatching answer. Expected %d, got %d
				`, input.ans, ans)
		}
	}
}
