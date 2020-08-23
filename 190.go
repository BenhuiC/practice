package partice

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		tmp := num & 1
		num = num >> 1
		result = result<<1 + tmp
	}
	return result
}
