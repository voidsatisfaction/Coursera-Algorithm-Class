package test

import (
	"testing"

	"../pkg/heap"
)

type DummyIntHeap []int

func (h DummyIntHeap) Len() int {
	return len(h)
}

func (h DummyIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h DummyIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DummyIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *DummyIntHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	top := old[n-1]
	*h = old[0 : n-1]
	return top
}

func TestHeap(t *testing.T) {
	h := &DummyIntHeap{}
	heap.Push(h, 1)
	heap.Push(h, -1)
	heap.Push(h, 3)
	v := heap.Pop(h)

	if v != -1 {
		t.Errorf("heap.Pop() method is not right")
		t.Errorf("Expect: %d, got: %d", -1, v)
	}

	h = &DummyIntHeap{1, 7, 9, 3, 2, 2, 5, 4, 6, 8}
	heap.Init(h)
	expected := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; h.Len() > 0; i++ {
		v := heap.Pop(h)
		if expected[i] != v {
			t.Errorf("Expected: %+v, got: %+v\n", expected, *h)
			break
		}
	}
}
