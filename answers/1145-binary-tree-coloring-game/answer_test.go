package answer1145

import (
	"leetcode/utils"
	"leetcode/utils/result"
	"leetcode/utils/structures"
	"testing"
)

func TestAnswer(t *testing.T) {
	n := 11
	tree := structures.NewTree(utils.Range(1, n))
	ans := btreeGameWinningMove(tree, n, 3)
	res := result.New(ans, true)
	if !res.Accepted() {
		t.Fatal(res)
	}
	t.Log(res)
}
