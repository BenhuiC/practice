package partice

func angleClock(hour int, minutes int) float64 {
	hourD := float64(hour)*30.0 + float64(minutes)/2.0
	minutesD := float64(minutes) * 6.0
	res1344 := minutesD - hourD
	if res1344 < 0 {
		res1344 = -res1344
	}
	if res1344 > 180 {
		res1344 = 360 - res1344
	}
	return res1344
}
