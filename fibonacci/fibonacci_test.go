package fibonacci

import "testing"

type testpair struct {
	val, fib int
}

var tests = []testpair{
	{-1, 0},
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{32, 2178309},
}

func TestFibonacci(t *testing.T) {
	for _, pair := range tests {
		v := CalFibonacci(pair.val)
		if v != pair.fib {
			t.Error(
				"For", pair.val,
				"expected", pair.fib,
				"got", v,
			)
		}
	}
}
