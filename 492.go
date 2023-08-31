package partice

import "math"

func constructRectangle(area int) []int {
	res492 := make([]int, 2)
	sq := int(math.Sqrt(float64(area)))
	l, w := sq, sq
	for {
		tmpArea := l * w
		if tmpArea == area {
			res492[0] = l
			res492[1] = w
			break
		}
		if tmpArea < area {
			l++
		} else {
			w--
		}
	}

	return res492
}
