package array

func minFlipsMonoIncr(s string) int {
	var dp0, dp1 int

	for _, c := range s {
		if c == '1' {
			dp1 = min(dp0, dp1)
			dp0 += 1
		} else {
			dp1 = 1 + min(dp0, dp1)
		}
	}

	return min(dp0, dp1)
}
