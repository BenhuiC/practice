package hot

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	find := func(ary []int, target int) int {
		l, r := 0, len(ary)-1
		if r < 0 {
			return -1
		}
		for l <= r {
			mid := (r-l)/2 + l
			if ary[mid] == target {
				return mid
			} else if ary[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return -1
	}

	for i := 0; i < n-2; i++ {
		if 0-nums[i] < nums[i] {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			if target > nums[n-1] {
				continue
			}
			if target < nums[j] {
				break
			}
			if k := find(nums[j+1:], target); k != -1 {
				res = append(res, []int{nums[i], nums[j], nums[j+k+1]})
			}
		}
	}

	return res
}
