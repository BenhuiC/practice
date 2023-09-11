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
	dp := make([][]bool, nlen)
	for i := 0; i < nlen; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < nlen; i++ {
		curNum := nums[i]
		for j := 1; j <= target; j++ {
			if j >= curNum {
				dp[i][j] = dp[i-1][j-curNum] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[nlen-1][target]
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
