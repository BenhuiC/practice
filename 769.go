package partice

func maxChunksToSorted(arr []int) int {
	var res769 int
	n := len(arr)
	for i := 0; i < n; {
		if arr[i] == i {
			res769++
			i++
			continue
		}
		mx, mi := arr[i], i
		var j int
		for j = i + 1; j < n; j++ {
			if arr[j] > mx {
				mx = arr[j]
			}
			if j-i == mx-mi {
				res769++
				break
			}
		}
		i = j + 1
	}
	return res769
}
