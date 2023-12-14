package offer

func canPlaceFlowers(flowerbed []int, n int) bool {
	ln := len(flowerbed)
	for i := 0; i < len(flowerbed); {
		if n == 0 {
			break
		}
		if flowerbed[i] == 1 {
			i += 2
			continue
		}
		if i > 0 && flowerbed[i-1] == 1 || i < ln-1 && flowerbed[i+1] == 1 {
			i++
		} else {
			flowerbed[i] = 1
			n--
			i += 2
		}

	}

	return n == 0
}
