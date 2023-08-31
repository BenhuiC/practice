package partice

import (
	"fmt"
	"math"
)

/*
一张 为期一天 的通行证售价为 costs[0] 美元；
一张 为期七天 的通行证售价为 costs[1] 美元；
一张 为期三十天 的通行证售价为 costs[2] 美元

可得
一周内如果天数大于等于costs[1]/costs[0]，则购买周票
一月内如果需要的周数大于等于costs[2]/costs[1]，则购买月票


问题；
可能存在日票比周票实惠，周票比月票实惠的情况（根据上述逻辑，这种情况不需要特殊考虑）
*/

/*
解决方案：
dp -> 从尾部开始dp[i]= min(dp[i+1]+costs[0],dp[i+7]+costs[1],dp[i+30]+costs[2])
贪心 -> 如何选择周期？
*/
func mincostTickets(days []int, costs []int) int {
	dpM := make([]int, 366)
	dayMap := make(map[int]bool)
	for _, v := range days {
		dayMap[v] = true
	}

	var f func(d int) int
	f = func(d int) int {
		if d >= 366 {
			return 0
		}
		if dpM[d] > 0 {
			return dpM[d]
		}
		if !dayMap[d] {
			dpM[d] = f(d + 1)
		} else {
			dpM[d] = Min(f(d+1)+costs[0], Min(f(d+7)+costs[1], f(d+30)+costs[2]))
		}
		return dpM[d]
	}

	return f(1)
}

func mincostTickets2(days []int, costs []int) int {
	dpM := make([]int, len(days))
	duration := []int{1, 7, 30}

	var dp func(d int) int
	dp = func(d int) int {
		if d >= len(days) {
			return 0
		}
		if dpM[d] > 0 {
			return dpM[d]
		}

		dpM[d] = math.MaxInt
		j := d
		for i := 0; i < len(duration); i++ {
			for ; j < len(days) && days[j] < days[d]+duration[i]; j++ {
			}
			dpM[d] = Min(dpM[d], dp(j)+costs[i])
		}

		return dpM[d]
	}
	res := dp(0)
	fmt.Println(dpM)
	return res
}
