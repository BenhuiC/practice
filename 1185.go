package partice

var dd =[]string{ "Friday", "Saturday","Sunday", "Monday", "Tuesday", "Wednesday", "Thursday"}
var mm=[]int{0,31,60,91,121,152,182,213,244,274,305,335,366}

// todo
func dayOfTheWeek(day int, month int, year int) string {
	days:=0
	for i:=1971;i<year;i++{
		days+=mm[12]
		if !(i%400==0)&&!(i%4==0&&i%100!=0){
			days-=1
		}
	}
	flag:=(year%400==0)||(year%4==0&&year%100!=0)
	days=mm[month-1]+day
	if !flag&&month>2{
		days--
	}

	return dd[days%7]
}