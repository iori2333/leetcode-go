package answer0001

import (
	"leetcode/utils/result"
	"testing"
)

func TestAnswer(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	res := result.NewSlice(ans, []int{0, 1})
	if !res.Accepted() {
		t.Fatal(res)
	}
	t.Log(res)
}
