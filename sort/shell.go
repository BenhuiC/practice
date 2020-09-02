package sort

func ShellSort(ary []int) {
	for gap := len(ary) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(ary); i++ {
			j := i - gap
			current := ary[i]
			for ; j >= 0 && ary[j] > current; j = j - gap {
				ary[j+gap] = ary[j]
			}
			ary[j+gap] = current
		}
	}
}
