package partice

func singleNumber(nums []int) int {
	var a int
	for _, v := range nums {
		a = a ^ v
	}
	return a
}
