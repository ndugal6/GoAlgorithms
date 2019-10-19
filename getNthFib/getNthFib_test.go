package algoExpert

import (
	"reflect"
	"testing"
)


func TestGetNthFib(t *testing.T) {
	cases := []struct {
		input          int
		expectedOutput int
	}{
		{
			6,
			5,
		},
		{
			0,
			0,
		},
		{
			2,
			1,
		},
		{
			10,
			34,
		},
	}
	for _, c := range cases {
		got := GetNthFib(c.input)
		if !reflect.DeepEqual(got, c.expectedOutput) {
			t.Errorf("Nth Fib Wrong - nthFib(%d) == %d, want %d", c.input, got, c.expectedOutput)
		}
	}
}
