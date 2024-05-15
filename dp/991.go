package dp

func brokenCalc(startValue int, target int) int {
	ans := 0
	for startValue < target {
		ans++
		if target&1 == 1 {
			target++
		} else {
			target /= 2
		}
	}

	return ans + startValue - target
}
