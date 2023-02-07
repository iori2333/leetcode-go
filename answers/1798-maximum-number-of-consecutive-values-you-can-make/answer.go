package answer1798

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	lowest := 1
	for _, coin := range coins {
		if coin > lowest {
			break
		}
		lowest += coin
	}
	return lowest
}
