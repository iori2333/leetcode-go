package answer0022

import (
	"leetcode/utils/result"
	"testing"
)

func TestAnswer(t *testing.T) {
	n := 3
	ans := generateParenthesis(n)
	gt := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

	res := result.NewSlice(gt, ans)
	if !res.Accepted() {
		t.Fatal(res)
	}
	t.Log(res)
}
