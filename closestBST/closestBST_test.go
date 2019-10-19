package algoExpert

import (
	"testing"
)

func TestFindClosestValue(t *testing.T) {
	cases := []struct {
		inputTree      BST
		targetValue    int
		expectedOutput int
	}{
		{d1_10, 10, 10},
		{d1_10, 5, 5},
		{d1_10, 1, 1},
		{d1_10, 5, 5},
		{d1_10, -10, 1},
		{d1_10, 50, 22},
		{d1_10, -50, 1},
		{d1_10, 6, 5},
		{d1_10, 16, 15},

		{leaf1, 100, 1},
		{leaf1, 1, 1},
		{leaf1, 0, 1},
	}
	for _, testCase := range cases {
		got := testCase.inputTree.FindClosestValue(testCase.targetValue)
		if got != testCase.expectedOutput {
			t.Errorf("Reverse(%d) == %d, want %d", testCase.targetValue, got, testCase.expectedOutput)
		}
	}
}
