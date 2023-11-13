package offer

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
