package partice

import "fmt"

func sortColors(nums []int) {
	l, f := 0, len(nums)-1
	for l < f {
		for nums[l] == 0 && l < f {
			l++
		}
		for nums[f] != 0 && l < f {
			f--
		}
		nums[l], nums[f] = nums[f], nums[l]
	}
	fmt.Println(nums)
	f = len(nums) - 1
	for l < f {
		for nums[l] == 1 && l < f {
			l++
		}
		for nums[f] != 1 && l < f {
			f--
		}
		nums[l], nums[f] = nums[f], nums[l]
	}
	fmt.Println(nums)
}
