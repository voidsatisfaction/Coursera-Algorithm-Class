package hash

import (
	"container/list"
)

type SCHash struct {
	M int

	Data []*list.List
}

func NewSeparateChainingHash() *SCHash {
	Data := make([]*list.List, M)
	for i := 0; i < M; i++ {
		Data[i] = list.New()
	}
	return &SCHash{
		M:    0,
		Data: Data,
	}
}

func (sch *SCHash) Put(hk HashKey, val interface{}) {
	i := (hk.hashCode() % M)

	for e := sch.Data[i].Front(); e != nil; e = e.Next() {
		if e.Value.(*node).key.equals(hk) {
			e.Value.(*node).val = val
			return
		}
	}

	newNode := &node{
		key: hk,
		val: val,
	}
	sch.Data[i].PushFront(newNode)
}

func (sch *SCHash) Get(hk HashKey) (interface{}, bool) {
	i := (hk.hashCode() % M)

	for e := sch.Data[i].Front(); e != nil; e = e.Next() {
		if e.Value.(*node).key.equals(hk) {
			return e.Value.(*node).val, true
		}
	}

	return nil, false
}
