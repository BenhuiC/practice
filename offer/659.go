package offer

func isPossible(nums []int) bool {
	left := make(map[int]int) // 数字剩余的个数
	for _, v := range nums {
		left[v]++
	}
	endCnt := make(map[int]int) // 以当前数字结尾的子序列个数
	for _, v := range nums {
		if left[v] == 0 {
			continue
		}
		if endCnt[v-1] > 0 { // 如果有以v-1为结尾的子序列，则该子序列可以以v结尾
			endCnt[v-1]--
			endCnt[v]++
			left[v]--
		} else if left[v+1] > 0 && left[v+2] > 0 {
			left[v]--
			left[v+1]--
			left[v+2]--
			endCnt[v+2]++
		} else {
			return false
		}
	}
	return true
}

// 超时
func isPossible2(nums []int) bool {
	nl := len(nums)
	if nl < 3 {
		return false
	}
	var dfs func(cur, curl int) bool
	dfs = func(cur, curl int) (res bool) {
		if curl == nl {
			return true
		}
		val := nums[cur]
		nums[cur] = -1
		defer func() {
			nums[cur] = val
		}()
		if curl < 3 {
			for i := cur + 1; i < nl; i++ {
				if nums[i] == -1 || nums[i] == val {
					continue
				}
				if nums[i]-val > 1 {
					break
				}
				return dfs(i, curl+1)
			}
			return false
		}
		r1 := true
		for i := 0; i < nl; i++ {
			if nums[i] != -1 {
				r1 = dfs(i, 1)
				break
			}
		}
		if r1 {
			return true
		}

		for i := cur + 1; i < nl; i++ {
			if nums[i] == -1 || nums[i] == val {
				continue
			}
			if nums[i]-val > 1 {
				break
			}
			return dfs(i, curl+1)
		}
		return false
	}

	return dfs(0, 1)
}
