package subsetsum

import (
	"fmt"
	"sort"
	"testing"
)

func TestSubsetSum(t *testing.T) {
	weights := []int{15, 22, 14, 26, 32, 9, 16, 1}
	target := 60 // 9 14 15 22
	sort.Ints(weights)

	total := 0
	for _, val := range weights {
		total += val
	}

	initialRes(len(weights))
	if weights[0] <= target && total >= target {
		subset_sum(weights, 0, 0, target)
	}

	for i, val := range *res {
		if val == 1 {
			fmt.Printf("%d ", weights[i])
		}
	}
	fmt.Println()
}

func initialRes(len int) {
	tmp := make([]int, len)
	res = &tmp
}
