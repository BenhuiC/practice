package partice

var presult [][]int

func permute(nums []int) [][]int {
	presult = make([][]int, 0)
	if len(nums) == 0 {
		return presult
	}
	a := make([]int, 0)
	pai(nums, a, 0)
	return presult
}

func pai(nums, ary []int, index int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, ary)
		presult = append(presult, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		ary = append(ary, nums[len(ary)])
		pai(nums, ary, len(ary))
		ary = ary[:len(ary)-1]
		nums[i], nums[index] = nums[index], nums[i]
	}
}
