package heap

import (
	"container/heap"
	"fmt"
	"testing"
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
	heap.Init(h)
	heap.Push(h, 1)
	heap.Push(h, 2)
	heap.Push(h, -1)
	fmt.Printf("%+v", h)
}
