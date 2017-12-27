package heap

import (
	"sort"
)

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func up(h Interface, k int) {
	for {
		p := (k - 1) / 2
		if p == k || !h.Less(k, p) {
			break
		}
		h.Swap(p, k)
		k = p
	}
}

func down(h Interface, k, n int) {
	for (2*k + 1) < n {
		j := 2*k + 1
		if j+1 < n && h.Less(j+1, j) {
			j++
		}
		if !h.Less(j, k) {
			break
		}
		h.Swap(k, j)
		k = j
	}
}
