package array

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	n := len(queries)
	if n == 0 {
		return nil
	}
	res := make([]int, 0, n)
	odd := 0
	for _, v := range nums {
		if v&1 == 0 {
			odd += v
		}
	}

	for _, q := range queries {
		val, idx := q[0], q[1]
		if val&1 == nums[idx]&1 {
			if nums[idx]&1 == 0 {
				odd += val
			} else {
				odd += val + nums[idx]
			}
		} else if nums[idx]&1 == 0 {
			odd -= nums[idx]
		}
		nums[idx] += val
		res = append(res, odd)
	}

	return res
}
