package mysort

func InsertSort(ary []int) {
	for i := 1; i < len(ary); i++ {
		j := i - 1
		current := ary[i]
		for ; j >= 0 && ary[j] > current; j-- {
			ary[j+1] = ary[j]
		}
		ary[j+1] = current
	}
}
