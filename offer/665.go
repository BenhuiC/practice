package offer

func checkPossibility(nums []int) bool {
	lessCnt := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x <= y {
			continue
		}
		lessCnt++
		if lessCnt > 1 {
			return false
		}
		if i > 0 && nums[i-1] > y {
			nums[i+1] = x
		} else {
			nums[i] = y
		}
	}

	return true
}

func checkPossibility2(nums []int) bool {
	stack := make([]int, 0)
	lessCnt := 0
	for _, v := range nums {
		if len(stack) == 0 || stack[len(stack)-1] <= v {
			stack = append(stack, v)
			continue
		}
		less := 0
		for i := len(stack) - 1; i >= 0; i-- {
			if less > 1 {
				break
			}
			if stack[i] > v {
				less++
			} else {
				break
			}
		}
		if less <= 1 {
			stack[len(stack)-1] = v
			stack = append(stack, v)
		} else {
			stack = append(stack, stack[len(stack)-1])
		}

		lessCnt++
		if lessCnt > 1 {
			return false
		}
	}
	return true
}
