package offer

// Boyer-Moore 投票算法
func majorityElement(nums []int) int {
	res := 0
	cnt := 0
	for _, v := range nums {
		if cnt == 0 {
			res = v
			cnt++
			continue
		}
		if res == v {
			cnt++
		} else {
			cnt--
		}
	}

	return res
}
