package partice

var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	if len(nums) <= 0 {
		return result
	}
	a := make([]int, 0)
	re(nums, a, 0)
	return result
}

func re(nums, a []int, n int) {
	tmp := make([]int, len(a))
	copy(tmp, a)
	result = append(result, tmp)
	for i := n; i < len(nums); i++ {
		a = append(a, nums[i])
		re(nums, a, i+1)
		a = a[:len(a)-1]
	}
}
