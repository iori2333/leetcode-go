package answer1604

import (
	"sort"
	"strconv"
)

func parseTime(keyTime string) int {
	hour, _ := strconv.Atoi(keyTime[:2])
	minute, _ := strconv.Atoi(keyTime[3:])
	return hour*60 + minute
}

func alertNames(keyName []string, keyTime []string) (res []string) {
	record := map[string][]int{}
	for i := range keyTime {
		record[keyName[i]] = append(record[keyName[i]], parseTime(keyTime[i]))
	}

	for name, times := range record {
		if len(times) < 3 {
			continue
		}
		sort.Ints(times)
		for i := 0; i < len(times)-2; i++ {
			if times[i+2]-times[i] <= 60 {
				res = append(res, name)
				break
			}
		}
	}

	sort.Strings(res)
	return
}
