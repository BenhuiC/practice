package hot

func sortColors(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	idx0, idx1 := 0, 0
	for i, v := range nums {
		if v == 0 {
			nums[i], nums[idx0] = nums[idx0], nums[i]
			if idx0 < idx1 {
				nums[idx1], nums[i] = nums[i], nums[idx1]
			}
			idx0++
			idx1++
		} else if v == 1 {
			nums[i], nums[idx1] = nums[idx1], nums[i]
			idx1++
		}
	}

}

func sortColors2(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	l, r := 0, n-1
	for l < r {
		for l < r && nums[l] != 2 {
			l++
		}
		for l < r && nums[r] != 0 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	l, r = 0, n-1
	for l < r {
		for l < r && nums[l] != 1 {
			l++
		}
		for l < r && nums[r] != 0 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	l, r = 0, n-1
	for l < r {
		for l < r && nums[l] != 2 {
			l++
		}
		for l < r && nums[r] != 1 {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
}
