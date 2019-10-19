package algoExpert

import (
	"reflect"
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	cases := []struct {
		inputNumbers   []int
		targetSum      int
		expectedOutput []int
	}{
		{
			[]int{3, 5, -4, 8, 11, 1, -1, 6},
			10,
			[]int{-1, 11},
		},
		{
			[]int{-1, 0, 1, 2, 3}, 3, []int{0, 1, 2, 3},
		},
		{
			[]int{1, 0, -1, 3, 2}, 3, []int{0, 1, 2, 3},
		},
		{
			[]int{-100, 1, 100, 0, -1, 3, -50, 2}, 3, []int{0, 1, 2, 3},
		},
		{
			[]int{1, -1}, 0, []int{-1, 1},
		},
		{
			[]int{-1, -2, -3, -4, -5, -6}, -11, []int{-6, -5},
		},
	}
	for _, c := range cases {
		got := TwoNumberSum(c.inputNumbers, c.targetSum)
		if !reflect.DeepEqual(got, c.expectedOutput) {
			t.Errorf("TwoNumberSum(%d, %d) == %d, want %d", c.inputNumbers, c.targetSum, got, c.expectedOutput)
		}
	}
}
