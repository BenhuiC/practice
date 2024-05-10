package array

func strWithout3a3b(a int, b int) string {
	res := ""
	for a > 0 || b > 0 {
		writeA := false
		n := len(res)
		if n >= 2 && res[n-1] == res[n-2] {
			if res[n-1] == 'b' {
				writeA = true
			}
		} else {
			if a >= b {
				writeA = true
			}
		}
		if writeA {
			res += "a"
			a--
		} else {
			res += "b"
			b--
		}
	}

	return res
}
