package offer

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mp := make(map[int]*Employee)
	for i, e := range employees {
		mp[e.Id] = employees[i]
	}
	res := 0
	ary := []int{id}
	for len(ary) != 0 {
		next := make([]int, 0)
		for _, v := range ary {
			if e, ok := mp[v]; ok {
				res += e.Importance
				next = append(next, e.Subordinates...)
			}
		}
		ary = next
	}
	return res
}
