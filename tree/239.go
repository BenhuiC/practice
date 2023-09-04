package tree

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	ary := make([]int, 0)
	l, r := 0, -1 // 加入第一个值后r才为0，所以初始值为-1
	for i, v := range nums {
		for len(ary) > 0 && nums[ary[len(ary)-1]] < v {
			ary = ary[:len(ary)-1]
		}
		ary = append(ary, i)
		r++

		if r-l+1 < k {
			continue
		} else {
			for len(ary) > 0 && ary[0] < l {
				ary = ary[1:]
			}
			res[l] = nums[ary[0]]
			l++
		}
	}

	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	for i, v := range nums {
		if i%k == 0 {
			prefixMax[i] = v
		} else {
			prefixMax[i] = Max(prefixMax[i-1], v)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffixMax[i] = nums[i]
		} else {
			suffixMax[i] = Max(suffixMax[i+1], nums[i])
		}
	}

	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = Max(suffixMax[i], prefixMax[i+k-1])
	}
	return ans
}
