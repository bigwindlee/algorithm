package rminvalid

func removeInvalidParenthesis(str string) []string {
	res := make([]string, 0)
	queue := []string{str}
	visit := make(map[string]int)
	level := false

	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:] // pop queue
		if isValid(el) {
			res = append(res, el)
			level = true
		}
		if level {
			// remove minimum number of parentheses, so stop at this level.
			continue
		}

		// Produce candidate and push into queue
		for i := 0; i < len(el); i++ {
			if el[i] == '(' || el[i] == ')' {
				sub := el[:i] + el[i+1:]
				if _, exist := visit[sub]; !exist {
					visit[sub]++
					queue = append(queue, sub) // push queue
				}
			}
		}
	}

	return res
}

func isValid(str string) bool {
	cnt := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			cnt++
		} else if str[i] == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}
