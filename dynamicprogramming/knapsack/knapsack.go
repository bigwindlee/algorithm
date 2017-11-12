package knapsack

import (
	"fmt"
	. "leetcomm"
)

func knapSack(W int, wt, val []int) int {
	options := []func(W int, wt, val []int) int{knapSack0, knapSack_r}
	return options[0](W, wt, val)
}

func knapSack_r(WT int, wt, val []int) int {
	fmt.Printf("%d, %v, %v\n", WT, wt, val)
	if len(wt) == 0 || WT == 0 {
		return 0
	}
	if wt[len(wt)-1] > WT {
		return knapSack_r(WT, wt[:len(wt)-1], val[:len(val)-1])
	}
	// a := val[len(val)-1] + knapSack_r(WT-wt[len(wt)-1], wt[:len(wt)-1], val[len(val)-1])
	// b := knapSack_r(WT, wt[:len(wt)-1], val[len(val)-1])
	return Max(val[len(val)-1]+knapSack_r(WT-wt[len(wt)-1], wt[:len(wt)-1], val[:len(val)-1]), knapSack_r(WT, wt[:len(wt)-1], val[:len(val)-1]))
}

func knapSack0(W int, wt, val []int) int {
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
			fmt.Printf("K[%d][%d] = %d\n", i, w, K[i][w])
			// fmt.Printf("\n** i=%d, w=%d **\n", i, w)
			// print2D(K)
		}
	}

	return K[n][W]
}

/*
func print2D(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%3d ", arr[i][j])
		}
		fmt.Println()
	}
}
*/
