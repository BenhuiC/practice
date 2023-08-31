package partice

import "math"

func maximumSum(arr []int) int {
	/*
		对包含第i个数的最大和区间如何判断？
		对于第i个数，其左区间的最大和为l，右区间的最大和为r
		如果l和r都为0，则表明左侧区间和右侧区间都为空，所以只能取当前值
		如果第i个数为负数，则当前区间的最大值为l+r
		如果为正数，则当前区间最大值为l+r+arr[i]

		对数组中的每个数根据上面的方式求最大值，即可得出答案

		如何求左区间的最大值？（右区间一样）
		第一个数的左区间和为0
		l := arr[i-1] + leftSumMax[i-1]
	*/
	res1186 := math.MinInt32
	leftSumMax := make([]int, len(arr))
	rightSumMax := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		l := arr[i-1] + leftSumMax[i-1]
		// 如果加上arr[i-1]的值为负数，则不需要加上其左边的值
		if l < 0 {
			l = 0
		}
		leftSumMax[i] = l

		j := len(arr) - i - 1
		r := arr[j+1] + rightSumMax[j+1]
		if r < 0 {
			r = 0
		}
		rightSumMax[j] = r
	}
	for i := 0; i < len(arr); i++ {
		l, r, cur := leftSumMax[i], rightSumMax[i], arr[i]
		if l == r && l == 0 {
			res1186 = Max(res1186, cur)
		} else {
			if cur > 0 {
				res1186 = Max(res1186, l+r+cur)
			} else {
				res1186 = Max(res1186, l+r)
			}
		}
	}

	return res1186
}

func maximumSum2(arr []int) int {
	/*
		动态规划
		dp0：以当前数字为结尾删除0个数字的区间最大和
		dp1：以当前数字为结尾删除1个数字的区间最大和
		状态转移方程：
		dp0=Max(dp0, 0) + arr[i] // 如果dp0<0，则不需要加上之前的区间，只保留当前值
		dp1=Max(dp0, dp1+arr[i]) // 删除当前值的最大值或在之前的区间删除过一次加当前值
	*/
	dp0, dp1, res := arr[0], 0, arr[0]
	for i := 1; i < len(arr); i++ {
		tmp0 := Max(dp0, 0) + arr[i]
		tmp1 := Max(dp0, dp1+arr[i])
		dp0, dp1 = tmp0, tmp1
		res = Max(res, Max(dp0, dp1))
	}

	return res
}
