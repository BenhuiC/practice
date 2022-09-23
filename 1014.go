package partice

func maxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0]+0
	for j := 1; j < len(values); j++ {
		ans = Max(ans, mx+values[j]-j)
		// 边遍历边维护
		mx = Max(mx, values[j]+j)
	}
	return ans
}
