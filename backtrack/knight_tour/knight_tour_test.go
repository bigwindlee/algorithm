package knighttour

import (
	"fmt"
	"testing"
)

func TestBacktrack(t *testing.T) {
	if solveKT() {
		fmt.Printf("%v\n", printSolution())
	} else {
		fmt.Println("Solution does not exist.")
	}

}

func printSolution() string {
	ba := make([]byte, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ba = append(ba, []byte(fmt.Sprintf("%2d ", sol[i][j]))...)
		}
		ba = append(ba, '\n')
	}
	return string(ba)
}
