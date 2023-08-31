package partice

func findComplement(num int) int {
	org := num
	numLen := 0
	for num != 0 {
		numLen++
		num = num >> 1
	}
	return org ^ (1<<numLen - 1)
}
