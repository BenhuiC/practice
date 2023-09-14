package dp

func rob213(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	var robRange func(a []int) int
	robRange = func(a []int) int {
		n := len(a)
		if n <= 0 {
			return 0
		}
		robMoney := a[0]
		notRobMoney := 0
		for i := 1; i < n; i++ {
			robMoney, notRobMoney = notRobMoney+a[i], Max(notRobMoney, robMoney)
		}

		return Max(robMoney, notRobMoney)
	}
	return Max(robRange(nums[:n-1]), robRange(nums[1:]))
}
