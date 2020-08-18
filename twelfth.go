package partice

/*
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

func intToRoman(num int) string {
	var s string
	s += getS(num / 1000)
	num = num % 1000
	s += getH(num / 100)
	num = num % 100
	s += getT(num / 10)
	num = num % 10
	s += getG(num)
	return s
}

func getS(n int) (r string) {
	for i := 0; i < n; i++ {
		r += "M"
	}
	return
}

func getH(n int) (r string) {
	switch {
	case n == 0:
		return
	case n == 9:
		r = "CM"
	case n == 4:
		r = "CD"
	case n >= 5:
		r = "D" + getH(n-5)
	default:
		for i := 0; i < n; i++ {
			r += "C"
		}
	}
	return
}

func getT(n int) (r string) {
	switch {
	case n == 0:
		return
	case n == 9:
		r = "XC"
	case n == 4:
		r = "XL"
	case n >= 5:
		r = "L" + getT(n-5)
	default:
		for i := 0; i < n; i++ {
			r += "X"
		}
	}
	return
}

func getG(n int) (r string) {
	switch {
	case n == 0:
		return
	case n == 9:
		r = "IX"
	case n == 4:
		r = "IV"
	case n >= 5:
		r = "V" + getG(n-5)
	default:
		for i := 0; i < n; i++ {
			r += "I"
		}
	}
	return
}

func intToRoman2(num int) string {
	var s string
	var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var labels = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(labels); i++ {
		for num >= values[i] {
			s += labels[i]
			num -= values[i]
		}
		if num <= 0 {
			break
		}
	}
	return s
}
