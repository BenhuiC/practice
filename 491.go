package partice

func findSubsequences(nums []int) [][]int {
	res491 := make([][]int, 0)
	ans := make([]int, 0)
	var backtrack func(i int)
	backtrack = func(index int) {
		if len(ans) > 1 {
			res491 = append(res491, append([]int{}, ans...))
		}
		m := make(map[int]bool)
		for i := index; i < len(nums); i++ {
			if m[nums[i]] {
				continue
			}
			if len(ans) == 0 || nums[i] >= ans[len(ans)-1] {
				m[nums[i]] = true
				ans = append(ans, nums[i])
				backtrack(i + 1)
				ans = ans[:len(ans)-1]
			}
		}
	}
	backtrack(0)
	return res491
}
