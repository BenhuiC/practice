package partice

func sequentialDigits(low int, high int) []int {
	res := make([]int, 0)
	var v int
	tmp := low
	for tmp != 0 {
		tmp /= 10
		v++
	}
	for {
		v1, v2 := genNu(v)
		if v1 > high {
			break
		}
		for i := 0; i <= 9-v; i++ {
			if v1 < low {
				v1 += v2
				continue
			}
			if v1 > high {
				break
			}
			res = append(res, v1)
			v1 += v2
		}
		v++
	}

	return res
}

func genNu(v int) (int, int) {
	var v1, v2 int
	for i := 1; i <= v; i++ {
		v1 = v1*10 + i
		v2 = v2*10 + 1
	}
	return v1, v2
}
