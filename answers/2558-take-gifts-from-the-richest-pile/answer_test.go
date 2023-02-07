package answer2558

import (
	"leetcode/utils/result"
	"testing"
)

func TestAnswer(t *testing.T) {
	gifts := []int{25, 64, 9, 4, 100}
	k := 4

	ans := pickGifts(gifts, k)
	res := result.New(ans, 29)
	if !res.Accepted() {
		t.Fatal(res)
	}
	t.Log(res)
}
