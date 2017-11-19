<h1><a href="http://www.geeksforgeeks.org/category/backtracking/">Backtracking</a></h1>


<a href="http://www.geeksforgeeks.org/backtracking-set-1-the-knights-tour-problem/">Backtracking | Set 1 (The Knight’s tour problem)</a><br />
<a href="http://www.geeksforgeeks.org/backttracking-set-4-subset-sum/">Backtracking | Set 4 (Subset Sum)</a><br />




<h2>回溯法总结</h2>
回溯法需要考虑的几个问题：
<ol>
<li>2个维度分别是什么；解空间是否可以包含重复的candidate，解空间中的candidates是否是有序的；</li>
<li>2个维度是否可以分别优化；
    <ul>
    <li>第1个维度用for循环实现，那么for循环的起点是否是变化的，是否可通过递归参数传递进来；</li>
    <li>第2个维度用递归实现，如何判定到达叶子节点（递归函数的出口）；如果是固定的深度则可通过depth递减来实现；如果是对candidates求和则可通过sum的递减来实现；</li>
    </ul>
</li>    
<li>isValid()函数的实现；尽量提前判断有效性；</li>
<li>dfs()搜索是对path的收集是否可以优化；</li>
</ol>
