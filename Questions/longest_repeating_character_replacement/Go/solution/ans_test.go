package solution_test

import (
	"replace/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		s   string
		k   int
		ans int
	}{
		{
			"ABAB",
			2,
			4,
		},
		{
			"ACBCCDEF",
			3,
			6,
		},
		{
			"XYYX",
			2,
			4,
		},
		{
			"AAABABB",
			1,
			5,
		},
	}

	for _, input := range inputs {
		ans := solution.CharacterReplacements(input.s, input.k)
		if ans != input.ans {
			t.Fatalf(`
				Mismatching answer. Expect %d, got %d
				`, input.ans, ans)
		}
	}
}
