package longestcommonsubsequence

import (
	// "fmt"
	. "leetcomm"
)

func lcs(X, Y string) int {
	options := []func(X, Y string) int{lcs0, lcs1}
	return options[1](X, Y)
}

/*----------------------------------------------------------------------------*/

/*
Naive recursion.
Time complexity of the naive recursive approach is O(2^n) in worst case and
worst case happens when all characters of X and Y mismatch i.e., length of LCS is 0.
*/
func lcs0(X, Y string) int {
	// fmt.Printf("#%4d\t%s\t%s\n", count, X, Y)
	if len(X) == 0 || len(Y) == 0 {
		return 0
	}
	if X[len(X)-1] == Y[len(Y)-1] {
		return 1 + lcs0(X[:len(X)-1], Y[:len(Y)-1])
	} else {
		return Max(lcs0(X, Y[:len(Y)-1]), lcs0(X[:len(X)-1], Y))
	}
}

/*----------------------------------------------------------------------------*/

/*
Time Complexity of the implementation is O(mn) which is much better than
the worst case time complexity of Naive Recursive implementation.
*/
func lcs1(X, Y string) int {
	L := initializeMatrix(len(X)+1, len(Y)+1) // L[i][j] contains length of LCS of X[0..i-1] and Y[0..j-1]
	for i := 0; i <= len(X); i++ {            // i stands for the variable range of the length of X, i.e. [0, len(X)]
		for j := 0; j <= len(Y); j++ { // j stands for the variable range of the length of Y, i.e. [0, len(Y)]
			if i == 0 || j == 0 {
				L[i][j] = 0
				// fmt.Printf("L[%d][%d] = %d\n", i, j, L[i][j])
			} else if X[i-1] == Y[j-1] { // i, j stands for the length of X, Y, so the last char is X[i-1], Y[j-1]
				L[i][j] = L[i-1][j-1] + 1
				// fmt.Printf("L[%d][%d] = %d\n", i, j, L[i][j])
			} else {
				L[i][j] = Max(L[i-1][j], L[i][j-1])
				// fmt.Printf("L[%d][%d] = %d\n", i, j, L[i][j])
			}
		}
	}
	return L[len(X)][len(Y)]
}

func initializeMatrix(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
