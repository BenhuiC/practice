package offer150

func convert(s string, numRows int) string {
	if len(s) <= numRows || 1 == numRows {
		return s
	}
	step := 2*numRows - 2
	level := 0
	res := ""
	for level = 0; level < numRows; level++ {
		lastStep := step - 2*level
		idx := level
		for idx < len(s) {
			if lastStep == 0 {
				lastStep = step
				continue
			}
			res += string(s[idx])
			idx += lastStep
			lastStep = step - lastStep
		}
	}
	return res
}
