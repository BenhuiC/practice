package partice

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}
	res := 1
	for i := 1; i < len(arr); {
		if arr[i] == arr[i-1] {
			i++
			continue
		}
		t := arr[i] > arr[i-1]
		j := i + 1
		for ; j < len(arr); j++ {
			t = !t
			if !bgNotEq(t, arr[j], arr[j-1]) {
				break
			}
		}
		res = max(res, j-i+1)
		i = j
	}

	return res
}

func bgNotEq(t bool, x, y int) bool {
	if x == y {
		return false
	}
	return x > y == t
}
