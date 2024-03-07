package hot

func findDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}
