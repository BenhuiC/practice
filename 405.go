package partice

import "math"

var numMap = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f"}

func toHex(num int) string {
	//2147483646
	//4294967295
	res405 := ""
	if res405 = numMap[num]; res405 != "" {
		return res405
	}
	if num < 0 {
		num = (num+math.MaxUint32)%math.MaxUint32 + 1
	}

	var c, y int
	for {
		c, y = num/16, num%16
		res405 = numMap[y] + res405
		if c == 0 {
			break
		}
		num = num / 16
	}
	return res405

}
