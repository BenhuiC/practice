package offer

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity > jug1Capacity+jug2Capacity {
		return false
	}
	if jug2Capacity == targetCapacity || jug1Capacity == targetCapacity {
		return true
	}
	l, g := jug1Capacity, jug2Capacity
	if l > g {
		l, g = g, l
	}
	z := l + g
	for z != 0 {
		if z == targetCapacity {
			return true
		}
		if z >= l {
			z = z - l
		} else {
			z = z + g
		}
	}
	return false
}
