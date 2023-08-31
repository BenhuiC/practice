package partice

import "sort"


// todo 未完成
func maxProfit(inventory []int, orders int) int {
	var rem int64=10e9+7
	var res int64
	if len(inventory)==0||orders<=0{
		return 0
	}
	sort.Ints(inventory)
	bg:=len(inventory)-1
	mx:=inventory[bg]
	for orders>0{
		var i int
		for i=bg;i>=0&&inventory[i]>=mx;i--{}
		od:=0
		mi:=0
		lw:=0
		if i>=0{
			mi=inventory[i]
			lw=i
		}
		od=(bg-i)*(mx-mi)
		if od<=orders{
			res+=int64((mx+mi+1)*(mx-mi)/2*(bg-lw))%rem
			mx=mi
		} else{
			d:=orders/(bg-lw)
			r:=orders%(bg-lw)
			for j:=d;j>0;j--{
				res+=int64((bg-lw)*mx)%rem
				mx--
			}
			res+=int64(r*mx)%rem
		}
		orders-=od
	}

	return int(res%rem)
}
