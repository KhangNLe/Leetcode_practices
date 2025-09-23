package solution_test

import (
	"testing"
	"versions/solution"
)

func TestSolution(t *testing.T) {
	inputs := []struct {
		v1     string
		v2     string
		expect int
	}{
		{
			"1.01",
			"1.001",
			0,
		}, {
			"1.0",
			"1.0.0.0",
			0,
		}, {
			"1.2",
			"1.10",
			-1,
		}, {
			"2.300000",
			"1.233",
			1,
		}, {
			"1.000000000001",
			"1.1",
			0,
		},
	}

	for _, input := range inputs {
		res := solution.CompareVersion(input.v1, input.v2)

		if res != input.expect {
			t.Fatalf(`
				Mismatching result, expect: %d, got %d\n
				`, input.expect, res)
		}
	}
}
