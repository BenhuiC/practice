package offer

import (
	"bytes"
	"fmt"
)

// timeout
func originalDigits2(s string) string {
	numAry := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	chartMap := make(map[string]map[byte]int)
	for _, v := range numAry {
		chartMap[v] = make(map[byte]int)
		for _, c := range v {
			chartMap[v][byte(c)]++
		}
	}
	chartNum := make([]int, 26)
	for _, v := range s {
		chartNum[v-'a']++
	}
	res := ""
	var backTrack func() bool
	backTrack = func() bool {
		fin := true
		for _, v := range chartNum {
			if v > 0 || v < 0 {
				fin = false
				break
			}
		}
		if fin {
			return true
		}
		for i, v := range numAry {
			has := true
			for c, nu := range chartMap[v] {
				if chartNum[c-'a'] < nu {
					has = false
					break
				}
			}
			if !has {
				continue
			}
			for c, nu := range chartMap[v] {
				chartNum[c-'a'] -= nu
			}
			res += fmt.Sprintf("%d", i)
			if backTrack() {
				return true
			}
			for c, nu := range chartMap[v] {
				chartNum[c-'a'] += nu
			}
			res = res[:len(res)-1]
		}
		return false
	}
	backTrack()
	return res
}

func originalDigits(s string) string {
	chartMap := make(map[byte]int)
	for _, v := range s {
		chartMap[byte(v)]++
	}
	cnt := [10]int{}
	cnt[0] = chartMap['z']
	cnt[2] = chartMap['w']
	cnt[4] = chartMap['u']
	cnt[6] = chartMap['x']
	cnt[8] = chartMap['g']

	cnt[3] = chartMap['h'] - cnt[8]
	cnt[5] = chartMap['f'] - cnt[4]
	cnt[7] = chartMap['s'] - cnt[6]

	cnt[1] = chartMap['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = chartMap['i'] - cnt[5] - cnt[6] - cnt[8]

	ans := []byte{}
	for i, v := range cnt {
		if v <= 0 {
			continue
		}
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, v)...)
	}

	return string(ans)
}
