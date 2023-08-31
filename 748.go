package partice

import "strings"

func shortestCompletingWord(licensePlate string, words []string) string {
	mp := make(map[byte]int)
	licensePlate = strings.ToLower(licensePlate)
	for i := range licensePlate {
		if licensePlate[i] == ' ' || licensePlate[i] >= '0' && licensePlate[i] <= '9' {
			continue
		}
		mp[licensePlate[i]]++
	}
	res := ""
	for idx, wd := range words {
		tmpMp := make(map[byte]int)
		wd = strings.ToLower(wd)
		for i := range wd {
			tmpMp[wd[i]]++
		}
		fg := true
		for k, v := range mp {
			if v > tmpMp[k] {
				fg = false
				break
			}
		}
		if fg {
			if res == "" {
				res = words[idx]
			} else if len(wd) < len(res) {
				res = words[idx]
			}
		}
	}

	return res
}
