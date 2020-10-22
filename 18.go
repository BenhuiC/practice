package partice

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i, j := 0, len(nums)-1; i < len(nums)-3 && i < j; {
		tmp := 0
		for x, y := i+1, j-1; x < y; {
			tmp = nums[i] + nums[j] + nums[x] + nums[y]
			if tmp == target {
				a := []int{nums[i], nums[x], nums[y], nums[j]}
				result = append(result, a)
				for x = x + 1; x < y && nums[x] == nums[x-1]; x++ {
				}
				for y = y - 1; x < y && nums[y+1] == nums[y]; y-- {
				}
			} else if tmp > target {
				y--
			} else {
				x++
			}
		}
		if nums[i]+nums[j]+nums[i+1]+nums[j-1] > target {
			j--
		} else {
			i++
		}
	}

	return result
}

var fourResult [][]int

func fourSum2(nums []int, target int) [][]int {
	fourResult = make([][]int, 0)
	if len(nums) < 4 {
		return fourResult
	}
	sort.Ints(nums)
	four(nums, make([]int, 0), 0, 0, target)
	return fourResult
}

func four(nums []int, ary []int, index, sum, target int) {
	if len(ary) == 4 {
		if sum == target {
			tmp := make([]int, 4)
			copy(tmp, ary)
			fourResult = append(fourResult, tmp)
			return
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if i != index && nums[i] == nums[i-1] {
			continue
		}
		ary = append(ary, nums[i])
		sum += nums[i]
		four(nums, ary, i+1, sum, target)
		ary = ary[:len(ary)-1]
		sum -= nums[i]
	}
}
