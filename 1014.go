package partice

func maxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0]+0
	for j := 1; j < len(values); j++ {
		ans = max(ans, mx+values[j]-j)
		// 边遍历边维护
		mx = max(mx, values[j]+j)
	}
	return ans
}
