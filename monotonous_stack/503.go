package monotonous_stack

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}

	ary := make([]int, 0)
	ary = append(ary, 0)
	idx := 1
	for len(ary) != 0 && idx < 2*len(nums) {
		i := idx % len(nums)
		if nums[ary[len(ary)-1]] > nums[i] {
			ary = append(ary, i)
			idx++
			continue
		}
		for len(ary) > 0 && nums[ary[len(ary)-1]] < nums[i] {
			res[ary[len(ary)-1]] = nums[i]
			ary = ary[:len(ary)-1]
		}
		ary = append(ary, i)
		idx++
	}

	return res
}
