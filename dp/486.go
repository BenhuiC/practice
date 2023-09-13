package dp

/*
dp1[i][j]表示为玩家1先手在nums[i,j]区间可以获得的最大积分
dp2[i][j]表示为玩家1后手在nums[i,j]区间可以获得的最大积分
dp1[i][j]=Max(nums[i]+dp2[i+1][j],nums[j]+dp2[i][j-1])
dp2[i][j]=Max(dp1[i+1][j],dp1[i][j-1])=Max()
dp1[i+1][j]=Max(nums[i+1]+dp2[i+2][j],nums[j]+dp2[i+1][j-1])
dp1[i][j-1]=Max(nums[i]+dp2[i+1][j-1],nums[j-1]+dp2[i][j-2])
*/
func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp1 := make([][]int, n)
	dp2 := make([][]int, n)
	for i := 0; i < n; i++ {
		dp1[i] = make([]int, n)
		dp2[i] = make([]int, n)
		dp1[i][i] = nums[i]
	}
	for i := n - 2; i >= 0; i-- {
		dp1[i][i+1] = Max(nums[i], nums[i+1])
		dp2[i][i+1] = Min(nums[i], nums[i+1])
		for j := i + 2; j < n; j++ {
			left := nums[i] + dp2[i+1][j]
			right := nums[j] + dp2[i][j-1]
			if left >= right {
				dp1[i][j] = left
				dp2[i][j] = dp1[i+1][j]
			} else {
				dp1[i][j] = right
				dp2[i][j] = dp1[i][j-1]
			}
		}
	}

	return dp1[0][n-1] >= dp2[0][n-1]
}
