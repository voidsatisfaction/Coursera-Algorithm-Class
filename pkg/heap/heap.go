package heap

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Interface) {}

func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0)
	return h.Pop()
}

func up(h Interface, k int) {
	for k > 0 && h.Less((k-1)/2, k) {
		h.Swap((k-1)/2, k)
		k = (k - 1) / 2
	}
}

func down(h Interface, k int) {
	for (2*k + 1) <= h.Len() {
		j := 2*k + 1
		if h.Less(j, j+1) {
			j++
		}
		if !h.Less(k, j) {
			break
		}
		h.Swap(k, j)
		k = j
	}
}
