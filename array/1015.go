package array

func smallestRepunitDivByK(k int) int {
	if k&1 == 0 || k%10 == 5 {
		return -1
	}
	res := 1
	lastMod := 1
	y := 0
	for {
		tmpY := lastMod % k
		y = (y + tmpY) % k
		if y == 0 {
			return res
		}
		lastMod = tmpY * 10
		res++
	}

	return -1
}
