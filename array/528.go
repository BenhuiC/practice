package array

import (
	"math/rand"
)

type Solution struct {
	data []int
	sv   int
}

func Constructor3(w []int) Solution {
	s := 0
	for i, v := range w {
		s += v
		w[i] = s
	}
	return Solution{
		data: w,
		sv:   s,
	}
}

func (this *Solution) PickIndex() int {
	res := 0
	r := rand.Intn(this.sv + 1)
	l, h := 0, len(this.data)-1
	var mid int
	for l <= h {
		mid = l + (h-l)/2
		if mid == 0 {
			if r <= this.data[mid] {
				res = mid
				break
			} else {
				l = mid + 1
				continue
			}
		}
		if this.data[mid-1] < r && this.data[mid] >= r {
			res = mid
			break
		} else if this.data[mid-1] >= r {
			h = mid - 1
		} else if this.data[mid] < r {
			l = mid + 1
		}
	}
	return res
}
