package offer

func validUtf8(data []int) bool {
	n := len(data)
	if n == 0 {
		return false
	}
	var byteLen, mask int
	headerIdx := 0
	for headerIdx < n {
		mask = 128
		byteLen = 0
		for mask&data[headerIdx] != 0 {
			mask = mask >> 1
			byteLen++
		}
		if byteLen == 0 {
			headerIdx++
			continue
		}
		if byteLen == 1 || byteLen > 4 {
			return false
		}
		if headerIdx+byteLen > n {
			return false
		}
		for i := headerIdx + 1; i < headerIdx+byteLen; i++ {
			if data[i]>>6 != 2 {
				return false
			}
		}
		headerIdx += byteLen
	}
	return true
}
