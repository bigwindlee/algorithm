package maze

import (
	"fmt"
	"testing"
)

func TestSolveMaze(t *testing.T) {
	if solveMazeUtil(0, 0) {
		printSolution()
	} else {
		fmt.Println("No solution exists.")
	}
}

func printSolution() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", sol[i][j])
		}
		fmt.Println()
	}
}
