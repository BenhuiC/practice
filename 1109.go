package partice

func corpFlightBookings(bookings [][]int, n int) []int {
	res1109:=make([]int,n)
	if n==0{
		return res1109
	}
	for _,v:=range bookings{
		for i:=v[0];i<=v[1];i++{
			res1109[i-1]+=v[2]
		}
	}
	return res1109
}
