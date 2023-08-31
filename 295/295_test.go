package _95

import (
	"testing"
)

/*
["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]
 */

func TestMedianFinder_AddNum(t *testing.T) {
	m:=Constructor()
	m.AddNum(-1)
	t.Log(m.FindMedian())
	m.AddNum(-2)
	t.Log(m.FindMedian())
	m.AddNum(-3)
	t.Log(m.FindMedian())
	m.AddNum(-4)
	t.Log(m.FindMedian())
	m.AddNum(-5)
	t.Log(m.FindMedian())
}

