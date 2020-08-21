package partice

func searchInsert(nums []int, target int) int {
	var result int
	if len(nums) == 0 {
		return result
	}
	tmp := []int{^int(^uint(0) >> 1)}
	tmp = append(tmp, nums...)
	tmp = append(tmp, ^int(^uint(0)>>1))
	s, e := 0, len(tmp)-1
	for s < e {
		t := (s + e) / 2
		if tmp[t] == target {
			return t - 1
		} else if tmp[t] > target {
			e = t
		} else {
			s = t
		}

		if s+1 == e {
			return s
		}
	}

	return e
}
