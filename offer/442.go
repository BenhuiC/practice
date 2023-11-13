package offer

func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := make([]int, 0)
	for i, v := range nums {
		if i+1 != v {
			res = append(res, v)
		}
	}
	return res
}

func findDuplicates2(nums []int) (ans []int) {
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		} else {
			ans = append(ans, x)
		}
	}
	return
}
