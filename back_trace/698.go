package back_trace

// not pass
func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	mx := 0
	for _, v := range nums {
		sum += v
		if v > mx {
			mx = v
		}
	}
	if sum%k != 0 {
		return false
	}
	part := sum / k
	if mx > part {
		return false
	}

	res := false
	n := len(nums)
	use := make([]bool, n)
	var backtrace func(sm, p int)
	backtrace = func(sm, p int) {
		if res == true {
			return
		}
		if sm == part {
			if p == k {
				res = true
				return
			}
			p++
			sm = 0
		}
		for i, v := range nums {
			if use[i] || sm+v > part {
				continue
			}
			use[i] = true
			backtrace(sm+v, p)
			use[i] = false
		}
	}
	backtrace(0, 1)

	return res
}
