package knapsack

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	testdata := []struct {
		val      []int
		wt       []int
		W        int
		expected int
	}{
		{[]int{60, 100, 120}, []int{10, 20, 30}, 50, 220},
	}

	for i := range testdata {
		el := testdata[i]
		actual := knapSack(el.W, el.wt, el.val)
		if actual != el.expected {
			t.Errorf("Case #%d\tactual: %d, expected: %d\n", i, actual, el.expected)
		}
	}
}
