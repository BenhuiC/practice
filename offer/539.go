package offer

import (
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) <= 1 || len(timePoints) > 1440 {
		return 0
	}
	sort.Strings(timePoints)
	cal := func(t string) int {
		ary := strings.Split(t, ":")
		if len(ary) != 2 {
			return 0
		}
		h, _ := strconv.Atoi(ary[0])
		m, _ := strconv.Atoi(ary[1])
		return h*60 + m
	}
	res := 24 * 60
	last := cal(timePoints[0])
	t0 := last
	for i := 1; i < len(timePoints); i++ {
		if timePoints[i] == timePoints[i-1] {
			return 0
		}
		tmp := cal(timePoints[i])
		if tmp-last < res {
			res = tmp - last
		}
		last = tmp
	}
	if t0+24*60-last < res {
		res = t0 + 24*60 - last
	}

	return res
}
