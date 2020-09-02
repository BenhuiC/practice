package sort

func MergeSort(ary []int) []int {
	if len(ary) <= 1 {
		return ary
	}
	mid := len(ary) / 2
	left := MergeSort(ary[:mid])
	right := MergeSort(ary[mid:])
	ary = merge(left, right)
	return ary
}

func merge(ary1 []int, ary2 []int) []int {
	if len(ary1) == 0 {
		return ary2
	}
	if len(ary2) == 0 {
		return ary1
	}
	result := make([]int, 0)
	m, n := 0, 0
	for {
		if m == len(ary1) || n == len(ary2) {
			break
		}
		if ary1[m] < ary2[n] {
			result = append(result, ary1[m])
			m++
		} else {
			result = append(result, ary2[n])
			n++
		}
	}
	if m == len(ary1) {
		result = append(result, ary2[n:]...)
	} else {
		result = append(result, ary1[m:]...)
	}
	return result
}
