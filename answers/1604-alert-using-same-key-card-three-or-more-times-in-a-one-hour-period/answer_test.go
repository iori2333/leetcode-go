package answer1604

import (
	"leetcode/utils/result"
	"testing"
)

func TestAnswer(t *testing.T) {
	keyName := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	keyTime := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}
	ans := alertNames(keyName, keyTime)
	res := result.NewSlice(ans, []string{"daniel"})
	if !res.Accepted() {
		t.Fatal(res)
	}
	t.Log(res)
}
