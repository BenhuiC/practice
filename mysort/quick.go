package mysort

func QuickSort(ary []int, left, right int) {
	if right-left <= 1 {
		return
	}
	t := partition(ary, left, right)
	QuickSort(ary, left, t-1)
	QuickSort(ary, t+1, right)
}

func partition(ary []int, left, right int) int {
	p := ary[left]
	l, h := left, right-1
	for l < h {
		for l < h && ary[h] > p {
			h--
		}
		if l < h {
			ary[l] = ary[h]
		}
		for l < h && ary[l] < p {
			l++
		}
		if l < h {
			ary[h] = ary[l]
		}
	}
	ary[l] = p
	return l
}
