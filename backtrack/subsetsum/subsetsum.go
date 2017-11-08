package subsetsum

/*
Backtracking | Set 4 (Subset Sum)
http://www.geeksforgeeks.org/backttracking-set-4-subset-sum/
*/

var res *[]int

func subset_sum(s []int, sum int, ite int, target_sum int) bool {
	if sum == target_sum {
		return true
	}

	for i := ite; i < len(s); i++ {
		if sum+s[i] <= target_sum {
			(*res)[i] = 1
			if subset_sum(s, sum+s[i], i+1, target_sum) {
				return true
			}
			(*res)[i] = 0 // BACKTRACK
		}
	}
	return false
}
