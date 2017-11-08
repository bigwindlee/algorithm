package knapsack

import (
	. "leetcomm"
)

func knapSack(W int, wt, val []int) int {
	n := len(val)
	K := make([][]int, n+1)
	for i := 0; i < len(K); i++ {
		K[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				K[i][w] = 0
			} else if wt[i-1] <= w {
				K[i][w] = Max(val[i-1]+K[i-1][w-wt[i-1]], K[i-1][w])
			} else {
				K[i][w] = K[i-1][w]
			}
		}
	}

	return K[n][W]
}
