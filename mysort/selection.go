package mysort

func SelectionSort(ary []int) {
	for i := 0; i < len(ary)-1; i++ {
		minIndex := i
		for j := i; j < len(ary); j++ {
			if ary[j] < ary[minIndex] {
				minIndex = j
			}
		}
		ary[i], ary[minIndex] = ary[minIndex], ary[i]
	}
}
