package array

func isLongPressedName(name string, typed string) bool {
	idx1, idx2 := 0, 0
	for idx1 < len(name) && idx2 < len(typed) {
		if name[idx1] != typed[idx2] {
			return false
		}

		tail1 := idx1 + 1
		for tail1 < len(name) && name[tail1] == name[idx1] {
			tail1++
		}
		tail2 := idx2 + 1
		for tail2 < len(typed) && typed[tail2] == typed[idx2] {
			tail2++
		}

		if tail1-idx1 > tail2-idx2 {
			return false
		}
		idx1 = tail1
		idx2 = tail2
	}

	return idx1 == len(name) && idx2 == len(typed)
}
