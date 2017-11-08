package nqueen

import (
	"fmt"
	"testing"
)

/* solution:

0 0 1 0
1 0 0 0
0 0 0 1
0 1 0 0

*/

func TestSolveNQ(t *testing.T) {
	if solveNQUtil(0) {
		fmt.Println()
		printSolution()
	} else {
		fmt.Println("Solution does not exist.")
	}
	fmt.Println()
}

func printSolution() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}
