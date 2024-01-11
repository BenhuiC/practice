package offer

func canTransform(start string, end string) bool {
	if start == end {
		return true
	}
	ls, le := len(start), len(end)
	if ls != le {
		return false
	}

	// todo

	return false
}
