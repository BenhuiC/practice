package partice

type minStackNode struct {
	rotDay   int
	appleNum int
}

type minStack []*minStackNode

func NewMinStk(n int) minStack {
	return make([]*minStackNode, 0, n)
}

func (m *minStack) Sink(n int) {
	for n*2+1 < len(*m) {
		if (*m)[n*2+1].rotDay < (*m)[n].rotDay || n*2+2 < len(*m) && (*m)[n*2+2].rotDay < (*m)[n].rotDay {
			if n*2+2 < len(*m) && (*m)[n*2+2].rotDay < (*m)[n*2+1].rotDay {
				(*m)[n*2+2], (*m)[n] = (*m)[n], (*m)[n*2+2]
				n = n*2 + 2
			} else {
				(*m)[n*2+1], (*m)[n] = (*m)[n], (*m)[n*2+1]
				n = n*2 + 1
			}
		} else {
			break
		}
	}
}

func (m *minStack) Up(n int) {
	for (*m)[n].rotDay < (*m)[(n-1)/2].rotDay && n >= 0 {
		(*m)[n], (*m)[(n-1)/2] = (*m)[(n-1)/2], (*m)[n]
		n = (n - 1) / 2
	}
}

func (m *minStack) Insert(node *minStackNode) {
	*m = append(*m, node)
	m.Up(len(*m) - 1)
}

func (m *minStack) DelMin() {
	l := len(*m)
	(*m)[0] = (*m)[l-1]
	*m = (*m)[:l-1]
	m.Sink(0)
}

func eatenApples(apples []int, days []int) int {
	res1705 := 0
	minStk := NewMinStk(len(apples))
	i := 0
	for {
		if i < len(apples) && apples[i] != 0 {
			minStk.Insert(&minStackNode{rotDay: i + days[i], appleNum: apples[i]})
		}
		if len(minStk) == 0 {
			if i >= len(apples) {
				break
			}
			i++
			continue
		}
		mi := minStk[0]
		if mi.rotDay <= i || mi.appleNum == 0 {
			minStk.DelMin()
			continue
		}
		res1705++
		i++
		mi.appleNum--
	}

	return res1705
}
