package answer0022

func generateParenthesis(n int) (res []string) {
	var backtrack func(self string, n, left, right int)
	backtrack = func(self string, n, left, right int) {
		if left == n && right == n {
			res = append(res, self)
			return
		}
		if left < n {
			backtrack(self+"(", n, left+1, right)
		}
		if right < left {
			backtrack(self+")", n, left, right+1)
		}
	}
	backtrack("", n, 0, 0)
	return
}
