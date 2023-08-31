package partice

import "time"

var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func dayOfYear(date string) int {
	t, _ := time.Parse("2006-01-02", date)
	var res1154 int
	for i := 0; i < int(t.Month())-1; i++ {
		res1154 += days[i]
	}
	if isRun(t.Year()) && t.Month() > 2 {
		res1154++
	}
	res1154 += t.Day()
	return res1154
}

func isRun(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%4 == 0 && y%100 == 0 {
		return true
	}
	return false
}
