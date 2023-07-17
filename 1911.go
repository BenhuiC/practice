package partice

func maxAlternatingSum(nums []int) int64 {
	var res1911 int64
	cur := 1
	for {
		// 找到波峰
		for cur < len(nums) && nums[cur] >= nums[cur-1] {
			cur++
		}
		l := cur - 1
		// 找到波谷
		for cur < len(nums) && nums[cur] <= nums[cur-1] {
			cur++
		}
		r := cur - 1
		// 波峰-波谷即为一对数字的最大值，不过最后一对数字只保留波峰
		if r == len(nums)-1 {
			res1911 += int64(nums[l])
			break
		} else {
			res1911 += int64(nums[l] - nums[r])
		}
	}
	return res1911
}
