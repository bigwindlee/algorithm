Remove Invalid Parentheses
http://www.geeksforgeeks.org/remove-invalid-parentheses/

An expression will be given which can contain open and close parentheses and optionally some characters, No other operator will be there in string. We need to remove minimum number of parentheses to make the input string valid. If more than one valid output are possible removing same number of parentheses then print all such output.
Examples:

Input  : str = “()())()” -
Output : ()()() (())()
There are two possible solutions
"()()()" and "(())()"

Input  : str = (v)())()
Output : (v)()()  (v())()

思路：
  1）注意"remove minimum number of parentheses"，找出最长的valid string即可，而不必深入其子集。在程序中使用level变量，一旦发现一个valid string，就停留在这个level，不再进入下一层level。
  2）注意queue的使用，把生成的candidate依次放入queue中以备检查是否是valid string。
  3）注意本题不能使用常规的backtrack解法，因为要求找出最长的valid string，常规的backtrack采用深度优先搜索，最先发现的valid string不一定是最长的！
  4）求最长的valid string，本质上是宽度优先搜索bfs，bfs算法一般使用queue来实现！
  5）回溯法似乎总是深度优先的，因为递归会一直调用自身直到找个一个叶子节点（递归出口）。
