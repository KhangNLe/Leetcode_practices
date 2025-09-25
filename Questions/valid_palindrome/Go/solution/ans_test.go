package solution_test

import (
	"testing"
	"valid/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		text   string
		result bool
	}{
		{
			"Was it a car or a cat I saw?",
			true,
		},
		{
			"tab a cat",
			false,
		},
		{
			"0P",
			false,
		},
		{
			".@$#@%aa",
			true,
		},
	}

	for _, input := range inputs {
		ans := solution.IsPalindrome(input.text)
		if ans != input.result {
			t.Fatalf(`
				Mismatching result. Expect %t, got %t
				`, input.result, ans)
		}
	}
}
