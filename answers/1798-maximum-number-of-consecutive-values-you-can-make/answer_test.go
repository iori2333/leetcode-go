package answer1798

import (
	"leetcode/utils/result"
	"testing"
)

func TestAnswer(t *testing.T) {
	coins := []int{1, 3}
	actual := getMaximumConsecutive(coins)
	res := result.New(actual, 2)
	if !res.Accepted() {
		t.Fatal(res)
	}
	t.Log(res)
}
