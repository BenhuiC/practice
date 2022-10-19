package partice

func countStudents(students []int, sandwiches []int) int {
	var num0, num1 int
	for _, v := range students {
		if v == 0 {
			num0++
		} else {
			num1++
		}
	}
	res := 0
	idx := 0
	for idx < len(sandwiches) {
		if sandwiches[idx] == 0 && num0 > 0 {
			num0--
			idx++
		} else if sandwiches[idx] == 1 && num1 > 0 {
			num1--
			idx++
		} else {
			res = num0 + num1
			break
		}
	}

	return res
}
