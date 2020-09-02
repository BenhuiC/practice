package sort

func BubbleSort(ary []int) {
	for i := len(ary) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if ary[j] > ary[j+1] {
				ary[j], ary[j+1] = ary[j+1], ary[j]
			}
		}
	}
}
