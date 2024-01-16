package offer

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	res := 0
	// last1: 上次值在[left,right]的idx，last2：上次值大于right的idx
	last1, last2 := -1, -1
	// 以i为结尾时子区间数量等于last1-last2
	for i, v := range nums {
		if v > right {
			last2 = i
			last1 = -1
			continue
		}
		if v >= left {
			last1 = i
		}
		if last1 != -1 {
			res += last1 - last2
		}
	}

	return res
}

func numSubarrayBoundedMax2(nums []int, left int, right int) int {
	// 最大值小于m的子数组数量
	cnt := func(m int) int {
		cur := 0
		res := 0
		for _, v := range nums {
			if v >= m {
				cur = 0
			} else {
				cur++
				res += cur
			}
		}
		return res
	}
	return cnt(right+1) - cnt(left)
}
