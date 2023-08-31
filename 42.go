package partice

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	var result int
	sta := make([]int, 0)
	var i int
	for ; height[i] == 0; i++ {
	}
	for j := i; j < len(height); {
		if len(sta) == 0 || height[j] <= height[sta[len(sta)-1]] {
			sta = append(sta, j)
			j++
		} else {
			var n int
			for n = j; n < len(height) && height[n] < height[sta[0]]; n++ {
			}
			if n == len(height) {
				sta = sta[1:]
				continue
			}
			var sum int
			for x := sta[0]; x < n; x++ {
				sum += height[x]
			}
			rtmp := Min(height[sta[0]], height[n])*(n-sta[0]) - sum
			result += rtmp
			sta = sta[:0]
			j = n
		}
	}
	return result
}
