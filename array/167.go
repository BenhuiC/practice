package array

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		for numbers[left]+numbers[right] > target {
			right--
		}
		if numbers[left]+numbers[right] == target {
			break
		}
		left++
	}
	return []int{left + 1, right + 1}
}
