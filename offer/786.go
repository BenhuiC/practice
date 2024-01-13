package offer

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	m := minheap{}
	m.push(&item{
		i:   0,
		j:   n - 1,
		Val: float32(arr[0]) / float32(arr[n-1]),
	})
	cnt := 0
	var res *item
	for cnt < k {
		res = m.pop()
		if res.i < res.j-1 {
			i, j := res.i+1, res.j
			m.push(&item{
				i:   i,
				j:   j,
				Val: float32(arr[i]) / float32(arr[j]),
			})
		}
		if res.i == 0 && res.j > 1 {
			i, j := 0, res.j-1
			m.push(&item{
				i:   i,
				j:   j,
				Val: float32(arr[i]) / float32(arr[j]),
			})
		}

		cnt++
	}

	return []int{arr[res.i], arr[res.j]}
}

type item struct {
	i, j int
	Val  float32
}

type minheap []*item

func (m *minheap) up() {
	for n := len(*m) - 1; n != 0; n = (n - 1) / 2 {
		r := (n - 1) / 2
		if (*m)[r].Val > (*m)[n].Val {
			(*m)[r], (*m)[n] = (*m)[n], (*m)[r]
			continue
		} else {
			break
		}
	}
}

func (m *minheap) down() {
	i := 0
	n := len(*m)
	for l := i*2 + 1; l < n; l = i*2 + 1 {
		if l < n-1 && (*m)[l+1].Val < (*m)[l].Val {
			l++
		}
		if (*m)[i].Val > (*m)[l].Val {
			(*m)[i], (*m)[l] = (*m)[l], (*m)[i]
			i = l
		} else {
			break
		}
	}
}

func (m *minheap) push(n *item) {
	*m = append(*m, n)
	m.up()
}

func (m *minheap) pop() *item {
	n := len(*m)
	if n == 0 {
		return nil
	}
	res := (*m)[0]
	(*m)[0] = (*m)[n-1]
	*m = (*m)[:n-1]
	m.down()

	return res
}
