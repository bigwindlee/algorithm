package permutation

import (
	. "leetcomm"
)

/*
Write a program to print all permutations of a given string
http://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
*/

func permute(nums []int) [][]int {
	options := []func(nums []int) [][]int{permute0, permute1}
	return options[1](nums)
}

///////////////////////////////////////////////////////////////////////////////

func permute1(nums []int) [][]int {
	res := make([][]int, 0)
	sol := make([]int, len(nums))
	visit := make([]bool, len(nums))
	backtrack1(&res, visit, sol, 0, nums)
	return res
}

func backtrack1(res *[][]int, visit []bool, sol []int, deep int, nums []int) {
	if deep == len(nums) {
		*res = append(*res, DupSlice(sol))
		return
	}

	for i, val := range nums {
		if !visit[i] {
			visit[i] = true
			sol[deep] = val
			backtrack1(res, visit, sol, deep+1, nums)
			visit[i] = false
		}
	}
}

///////////////////////////////////////////////////////////////////////////////

func permute0(nums []int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, 0)
	backtrack0(&res, &temp, nums)
	return res
}

func backtrack0(res *[][]int, temp *[]int, nums []int) {
	if len(*temp) == len(nums) {
		// Must duplicate the slice before append
		*res = append(*res, DupSlice(*temp))
	} else {
		for i := 0; i < len(nums); i++ {
			if isValid(*temp, nums[i]) {
				*temp = append(*temp, nums[i])
				backtrack0(res, temp, nums)
				*temp = (*temp)[:len(*temp)-1]
			}
		}
	}
}

/* Returns true when n is not contained in list */
func isValid(list []int, n int) bool {
	for _, val := range list {
		if val == n {
			return false
		}
	}
	return true
}
