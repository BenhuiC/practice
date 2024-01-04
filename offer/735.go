package offer

func asteroidCollision(asteroids []int) []int {
	st := make([]int, 0)
	for _, a := range asteroids {
		if len(st) == 0 || a > 0 {
			st = append(st, a)
			continue
		}
		for len(st) > 0 && st[len(st)-1] > 0 && st[len(st)-1] < (-a) {
			st = st[:len(st)-1]
		}
		if len(st) > 0 && st[len(st)-1] == -a {
			st = st[:len(st)-1]
		} else if len(st) == 0 || st[len(st)-1] < 0 {
			st = append(st, a)
		}
	}

	return st
}
