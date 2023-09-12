package dp

func canPartition(nums []int) bool {
	nlen := len(nums)
	maxNu := 0
	sum := 0
	for _, v := range nums {
		if v > maxNu {
			maxNu = v
		}
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	if maxNu > target {
		return false
	}
	if maxNu == target {
		return true
	}
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 0; i < nlen; i++ {
		curNum := nums[i]
		for j := target; j >= curNum; j-- {
			dp[j] = dp[j-curNum] || dp[j]
		}
	}

	return dp[target]
}

// 回溯法会超时
func canPartition2(nums []int) bool {
	nlen := len(nums)
	maxNu := 0
	sum := 0
	for _, v := range nums {
		if v > maxNu {
			maxNu = v
		}
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	if maxNu > target {
		return false
	}
	if maxNu == target {
		return true
	}

	s := 0
	var f func(i int) bool
	f = func(idx int) bool {
		if s == target {
			return true
		}
		if s > target {
			return false
		}
		if idx >= nlen {
			return false
		}

		for j := idx; j < nlen; j++ {
			s += nums[j]
			if f(j + 1) {
				return true
			}
			s -= nums[j]
		}
		return false
	}

	return f(0)
}
