package partice

func prevPermOpt1(arr []int) []int {
	var i, j int
	findI := false
	for x := len(arr) - 1; x > 0; x-- {
		if !findI && arr[x-1] > arr[x] {
			i = x - 1
			findI = true
			break
		}
	}
	if !findI {
		return arr
	}
	j = i + 1
	for z := i + 2; z < len(arr); z++ {
		if arr[z] > arr[j] && arr[z] < arr[i] {
			j = z
		}
	}

	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
