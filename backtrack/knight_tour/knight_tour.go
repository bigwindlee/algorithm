package knighttour

/*
Backtracking | Set 1 (The Knight’s tour problem)
http://www.geeksforgeeks.org/backtracking-set-1-the-knights-tour-problem/
*/

const (
	N = 8
)

var sol [N][N]int
var xMove []int = []int{2, 1, -1, -2, -2, -1, 1, 2}
var yMove []int = []int{1, 2, 2, 1, -1, -2, -2, -1}

///////////////////////////////////////////////////////////////////////////////

func solveKT() bool {
	initialSolution()
	sol[0][0] = 0
	return backtrack(0, 0, 1)
}

func backtrack(x int, y int, movei int) bool {
	if movei == N*N { // start FROM 0 TO (N*N-1)
		return true
	}

	// Try all next moves from the current coordinate x, y
	for k := 0; k < N; k++ {
		next_x := x + xMove[k]
		next_y := y + yMove[k]
		if isValid(next_x, next_y) {
			sol[next_x][next_y] = movei
			if backtrack(next_x, next_y, movei+1) {
				return true
			}
			sol[next_x][next_y] = -1 // backtracking
		}
	}

	return false
}

func isValid(x, y int) bool {
	return x >= 0 && x < N && y >= 0 && y < N && sol[x][y] == -1
}

func initialSolution() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sol[i][j] = -1
		}
	}
}
