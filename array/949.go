package array

import (
	"fmt"
)

func largestTimeFromDigits(arr []int) string {
	if len(arr) != 4 {
		return ""
	}

	findHour := func(x, y int) int {
		a, b := x*10+y, y*10+x
		if a < 24 && b < 24 {
			return max(a, b)
		} else if a < 24 {
			return a
		} else if b < 24 {
			return b
		}
		return -1
	}
	findMin := func(x, y int) int {
		a, b := x*10+y, y*10+x
		if a < 60 && b < 60 {
			return max(a, b)
		} else if a < 60 {
			return a
		} else if b < 60 {
			return b
		}
		return -1
	}

	hour, minute := -1, -1

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			a, b := arr[i], arr[j]
			tmpHour := findHour(a, b)
			if tmpHour >= 0 && tmpHour < hour {
				continue
			}

			arr[i], arr[j] = -1, -1
			x, y := -1, -1
			for _, v := range arr {
				if v < 0 {
					continue
				}
				if x == -1 {
					x = v
				} else {
					y = v
				}
			}
			arr[i], arr[j] = a, b
			tmpMin := findMin(x, y)
			if tmpMin < 0 {
				continue
			} else if tmpHour == hour {
				minute = max(minute, tmpMin)
			} else if tmpHour > hour {
				hour = tmpHour
				minute = tmpMin
			}
		}
	}
	if hour == -1 || minute == -1 {
		return ""
	}

	return fmt.Sprintf("%02d:%02d", hour, minute)
}
