package heap

import "sort"

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
	lastIndex := h.Len() - 1
	h.Swap(0, lastIndex)
	down(h, 0, lastIndex)
	return h.Pop()
}

func up(h Interface, i int) {
	for {
		p := (i - 1) / 2
		if i == 0 || h.Less(p, i) {
			break
		}
		h.Swap(i, p)
		i = p
	}
}

func down(h Interface, i, n int) {
	for {
		l := 2*i + 1
		if l >= n {
			break
		}
		if l+1 < n && h.Less(l+1, l) {
			l++
		}
		if !h.Less(l, i) {
			break
		}
		h.Swap(i, l)
		i = l
	}
}
