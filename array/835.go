package array

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)
	if n == 0 {
		return 0
	}

	ary := make([][]int, 2*n+1)
	for i := range ary {
		ary[i] = make([]int, 2*n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if img1[i][j] != 1 {
				continue
			}
			for i2 := 0; i2 < n; i2++ {
				for j2 := 0; j2 < n; j2++ {
					if img2[i2][j2] != 1 {
						continue
					}
					ary[i-i2+n][j-j2+n]++
				}
			}
		}
	}

	res := 0
	for _, c := range ary {
		for _, v := range c {
			res = max(res, v)
		}
	}

	return res
}
