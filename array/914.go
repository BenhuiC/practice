package array

func hasGroupsSizeX(deck []int) bool {
	count := make(map[int]int)
	for _, x := range deck {
		count[x]++
	}
	N := len(deck)
	for X := 2; X <= N; X++ {
		if N%X == 0 {
			valid := true
			for _, v := range count {
				if v%X != 0 {
					valid = false
					break
				}
			}
			if valid {
				return true
			}
		}
	}
	return false
}
