package offer

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	nagivate := false
	if num < 0 {
		nagivate = true
		num = -num
	}
	res := make([]byte, 0)
	for num != 0 {
		res = append([]byte{'0' + byte(num%7)}, res...)
		num /= 7
	}
	if nagivate {
		res = append([]byte{'-'}, res...)
	}
	return string(res)
}
