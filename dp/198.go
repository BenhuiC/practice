package dp

/*
robMoney[i]表示第i天抢钱的最大积分
notRobMoney[i]表示第i天不抢的最大积分

如果第i天不抢，则第i的的最大值为Max(robMoney[i-1],notRobMoney[i-1])，因为第i-1天可以抢也可以不抢
如果第i天抢，则第i-1天必须不抢，所以robMoney[i]=notRobMoney[i-1]+nums[i]
*/
func rob(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	robMoney := nums[0]
	notRobMoney := 0
	for i := 1; i < n; i++ {
		robMoney, notRobMoney = notRobMoney+nums[i], Max(notRobMoney, robMoney)
	}

	return Max(robMoney, notRobMoney)
}
