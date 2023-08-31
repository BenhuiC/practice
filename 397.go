package partice

func integerReplacement(n int) int {
	var res397 int
	for n != 1 {
		if n%2 == 0 {
			n /= 2
			res397++
		} else if n%4 == 1 {
			n = (n - 1) / 2
			res397 += 2
		} else if n%4 == 3 {
			if n == 3 {
				res397 += 2
				n = 1
				continue
			}
			n = (n + 1) / 2
			res397 += 2
		}
	}
	return res397
}
