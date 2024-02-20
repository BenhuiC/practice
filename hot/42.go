package hot

func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = height[i]
		if i > 0 && left[i-1] > left[i] {
			left[i] = left[i-1]
		}
	}
	for j := n - 1; j >= 0; j-- {
		right[j] = height[j]
		if j < n-1 && right[j+1] > right[j] {
			right[j] = right[j+1]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res += min(left[i], right[i]) - height[i]
	}

	return res
}

func trap2(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	sk := make([]int, 0)
	res := 0
	for r, h := range height {
		if len(sk) == 0 || height[sk[len(sk)-1]] > h {
			sk = append(sk, r)
			continue
		}
		for len(sk) > 0 && height[sk[len(sk)-1]] <= h {
			l := sk[len(sk)-1]
			sk = sk[:len(sk)-1]
			if len(sk) > 0 {
				t := sk[len(sk)-1]
				res += (r - t - 1) * (min(height[t], h) - height[l])
			}
		}
		sk = append(sk, r)
	}
	return res
}
