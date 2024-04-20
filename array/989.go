package array

func addToArrayForm(num []int, k int) []int {
	num2 := make([]int, 0)
	for k != 0 {
		num2 = append([]int{k % 10}, num2...)
		k /= 10
	}
	n1, n2 := len(num)-1, len(num2)-1
	y := 0
	res := make([]int, 0)
	for n1 >= 0 || n2 >= 0 {
		val := y
		if n1 >= 0 {
			val += num[n1]
			n1--
		}
		if n2 >= 0 {
			val += num2[n2]
			n2--
		}
		y = val / 10
		res = append([]int{val % 10}, res...)
	}
	if y > 0 {
		res = append([]int{y}, res...)
	}
	return res
}
