package dp

/*
nums的和为sum
为负数的部分和为neg
target=(sum-neg)-neg=sum-2*neg
target=sum-2*neg
neg=(sum-target)/2
所以需要求nums中和为neg的组合

dp[i][j]表示为前i个数的和为j的组合个数
对于当前的数x
dp[i][j]=dp[i-1][j]+dp[i-1][j-x]
由于处理当前层时只需要上一层，所以可以压缩dp数组
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < target {
		return 0
	}
	if (sum-target)&1 == 1 {
		return 0
	}
	neg := (sum - target) / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, n := range nums {
		for i := neg; i >= n; i-- {
			dp[i] = dp[i] + dp[i-n]

		}
	}

	return dp[neg]
}

func findTargetSumWays2(nums []int, target int) int {
	res := 0
	var backtrack func(idx, sum int)
	backtrack = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				res++
			}
			return
		}
		backtrack(idx+1, sum+nums[idx])
		backtrack(idx+1, sum-nums[idx])
	}
	backtrack(0, 0)
	return res
}
