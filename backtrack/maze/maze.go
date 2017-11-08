package maze

/*
Backtracking | Set 2 (Rat in a Maze)
http://www.geeksforgeeks.org/backttracking-set-2-rat-in-a-maze/
*/

const (
	N = 4
)

var maze = [][]int{
	[]int{1, 0, 0, 0},
	[]int{1, 1, 0, 1},
	[]int{0, 1, 0, 0},
	[]int{1, 1, 1, 1},
}

var sol = [][]int{
	[]int{0, 0, 0, 0},
	[]int{0, 0, 0, 0},
	[]int{0, 0, 0, 0},
	[]int{0, 0, 0, 0},
}

/////////////////////////////////////////////////////////////////////////////

/* A recursive utility function to solve Maze problem */
func solveMazeUtil(x, y int) bool {
	// if (x,y is goal) return true
	if x == N-1 && y == N-1 {
		sol[x][y] = 1
		return true
	}

	// Check if maze[x][y] is valid
	if isValidMove(x, y) {
		// mark x,y as part of solution path
		sol[x][y] = 1

		/* Move forward in x direction */
		if solveMazeUtil(x+1, y) {
			return true
		}

		/* If moving in x direction doesn't give solution then
		   Move down in y direction  */
		if solveMazeUtil(x, y+1) {
			return true
		}

		/* If none of the above movements work then BACKTRACK:
		   unmark x,y as part of solution path */
		sol[x][y] = 0
		return false
	}
	return false
}

/* A utility function to check if x,y is valid index for N*N maze */
func isValidMove(x, y int) bool {
	// if (x,y outside maze) return false
	return x >= 0 && x < N && y >= 0 && y < N && maze[x][y] == 1
}
