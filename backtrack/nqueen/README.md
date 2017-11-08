思路：

The idea is to place queens one by one in different columns, starting from the leftmost column. When we place a queen in a column, we check for clashes with already placed queens. In the current column, if we find a row for which there is no clash, we mark this row and column as part of the solution. If we do not find such a row due to clashes then we backtrack and return false.

1) Start in the leftmost column
2) If all queens are placed
    return true
3) Try all rows in the current column.  Do following for every tried row.
    a) If the queen can be placed safely in this row then mark this [row,
        column] as part of the solution and recursively check if placing  
        queen here leads to a solution.
    b) If placing queen in [row, column] leads to a solution then return
        true.
    c) If placing queen doesn't lead to a solution then umark this [row,
        column] (Backtrack) and go to step (a) to try other rows.
3) If all rows have been tried and nothing worked, return false to trigger


[row, col] mark，试探一个点为solution
(row, col) mark一个试探点之后，进入下一层递归的for循环，寻找下一个试探点；
{row, col} 试探失败之后的回退，clear试探点，并沿row方向寻找下一个试探点；

|-----[0, 0] // mark
|     (0, 1)
|     (1, 1)
|   |-[2, 1] // mark
|   | (0, 2)
|   | (1, 2)
|   | (2, 2)
|   | (3, 2)
|   |-{2, 1} // backtrack
| |---[3, 1] // mark
| |   (0, 2)
| | |-[1, 2] // mark
| | | (0, 3)
| | | (1, 3)
| | | (2, 3)
| | | (3, 3)
| | |-{1, 2} // backtrack
| |   (2, 2)
| |   (3, 2)
| |---{3, 1} // backtrack
|-----{0, 0} // backtrack
      [1, 0] // mark
      (0, 1)
      (1, 1)
      (2, 1)
      [3, 1] // mark
      [0, 2] // mark
      (0, 3)
      (1, 3)
      [2, 3] // mark done! Found a solution!
