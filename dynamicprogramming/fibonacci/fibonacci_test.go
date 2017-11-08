package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	testdata := [][]int{
		[]int{2, 1},
		[]int{3, 2},
		[]int{4, 3},
		[]int{5, 5},
		[]int{6, 8},
		[]int{42, 267914296},
	}

	for i := range testdata {
		actual := fibonacci(testdata[i][0])
		if testdata[i][1] != actual {
			t.Errorf("Case #%d\tactual: %d, expected: %d\n", i, actual, testdata[i][1])
		}
	}
}
