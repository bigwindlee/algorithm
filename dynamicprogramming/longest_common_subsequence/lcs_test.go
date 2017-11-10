package longestcommonsubsequence

import (
	"testing"
)

func TestLcs(t *testing.T) {
	testdata := []struct {
		X        string
		Y        string
		expected int
	}{
		{"AGGTAB", "GXTXAYB", 4},
		{"AGGTABAGGTAB", "GXTXAYBGXTXAYB", 8},
		{"AGGTABAGGTABAGGTAB", "GXTXAYBGXTXAYBGXTXAYB", 12},
	}

	for i := range testdata {
		el := testdata[i]
		actual := lcs(el.X, el.Y)
		if actual != el.expected {
			t.Errorf("Case #%d\tactual: %d, expected: %d\n", i, actual, el.expected)
		}
	}
}
