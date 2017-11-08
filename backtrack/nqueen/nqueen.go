package nqueen

const (
	N = 4
)

var board = [][]int{
	[]int{0, 0, 0, 0},
	[]int{0, 0, 0, 0},
	[]int{0, 0, 0, 0},
	[]int{0, 0, 0, 0},
}

/* A recursive utility function to solve N
   Queen problem */
func solveNQUtil(col int) bool {
	/* base case: If all queens are placed
	   then return true */
	if col >= N {
		return true
	}

	/* Consider this column and try placing
	   this queen in all rows one by one */
	for i := 0; i < N; i++ {
		/* Check if queen can be placed on
		   board[i][col] */
		if isSafe(i, col) {
			/* Place this queen in board[i][col] */
			board[i][col] = 1

			/* recur to place rest of the queens */
			if solveNQUtil(col + 1) {
				return true
			}

			/* If placing queen in board[i][col]
			   doesn't lead to a solution, then
			   remove queen from board[i][col] */
			board[i][col] = 0 // BACKTRACK
		}
	}

	/* If queen can not be place in any row in
	   this colum col  then return false */
	return false
}

/* A utility function to check if a queen can
   be placed on board[row][col]. Note that this
   function is called when "col" queens are
   already placed in columns from 0 to col -1.
   So we need to check only left side for
   attacking queens */
func isSafe(row, col int) bool {
	i, j := 0, 0

	/* Check this row on left side */
	for i = 0; i < col; i++ {
		if board[row][i] == 1 {
			return false
		}
	}

	/* Check upper diagonal on left side */
	for i, j = row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	/* Check lower diagonal on left side */
	for i, j = row, col; j >= 0 && i < N; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}
