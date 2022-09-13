package partice

func minOperations(logs []string) int {
	res := 0
	for _, v := range logs {
		switch v {
		case "../":
			if res > 0 {
				res--
			}
		case "./":
			// do nothing
		default:
			res++
		}
	}

	return res
}
