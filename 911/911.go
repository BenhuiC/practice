package _11

type TopVotedCandidate struct {
	q       []int
	persons []int
	times   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	q := make([]int, len(persons))
	q[0] = persons[0]
	mq := make(map[int]int)
	mq[persons[0]] = 1
	for i := 1; i < len(times); i++ {
		mq[persons[i]]++
		if mq[persons[i]] >= mq[q[i-1]] {
			q[i] = persons[i]
		} else {
			q[i] = q[i-1]
		}
	}
	return TopVotedCandidate{
		q:       q,
		persons: persons,
		times:   times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	target := 0
	for target < len(this.times) && this.times[target] <= t {
		target++
	}
	if target != 0 {
		target--
	}
	return this.q[target]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
