package partice

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	res := 0
	left, right := capacityA, capacityB
	leftIdx, rightIdx := 0, len(plants)-1
	for leftIdx <= rightIdx {
		if leftIdx == rightIdx {
			if left == right && left < plants[leftIdx] {
				res++
			} else if left > right && left < plants[leftIdx] {
				res++
			} else if right > left && right < plants[leftIdx] {
				res++
			}
			break
		}
		if left < plants[leftIdx] {
			left = capacityA - plants[leftIdx]
			res++
		} else {
			left -= plants[leftIdx]
		}
		if right < plants[rightIdx] {
			right = capacityB - plants[rightIdx]
			res++
		} else {
			right -= plants[rightIdx]
		}
		leftIdx++
		rightIdx--
	}
	return res
}
