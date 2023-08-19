package partice

func partitionDisjoint(nums []int) int {
	res := 0
	n := len(nums)
	ary := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			ary[i] = nums[i]
			continue
		}
		if nums[i] < ary[i+1] {
			ary[i] = nums[i]
		} else {
			ary[i] = ary[i+1]
		}
	}
	leftMax := 0
	for i := 0; i < n-1; i++ {
		if nums[i] > leftMax {
			leftMax = nums[i]
		}
		if leftMax < ary[i+1] {
			res = i + 1
			break
		}
	}
	return res
}
