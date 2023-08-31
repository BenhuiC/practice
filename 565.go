package partice

import "fmt"

func arrayNesting(nums []int) int {
	var res565 int
	var tmplen, last int
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			continue
		}
		tmplen, last = 0, i
		for nums[last] != -1 {
			fmt.Print(last)
			tmplen++
			res565 = max(res565, tmplen)
			tmplast := last
			last = nums[last]
			nums[tmplast] = -1
		}
		fmt.Println()
	}

	return res565
}
