package offer

func findTheArrayConcVal(nums []int) int64 {
	l, r := 0, len(nums)-1
	var res int64
	for l <= r {
		if l == r {
			res += int64(nums[l])
		} else {
			lv, rv := nums[l], nums[r]
			for rv > 0 {
				lv *= 10
				rv /= 10
			}
			res += int64(lv + nums[r])
		}
		l++
		r--
	}

	return res
}
