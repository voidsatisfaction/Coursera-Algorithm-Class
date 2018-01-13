package hash

import (
	"container/list"
)

type node struct {
	key Key

	val interface{}
}

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

func (sch *SCHash) Put(hk Key, val interface{}) {
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

func (sch *SCHash) Get(hk Key) (interface{}, bool) {
	i := (hk.hashCode() % M)

	for e := sch.Data[i].Front(); e != nil; e = e.Next() {
		if e.Value.(*node).key.equals(hk) {
			return e.Value.(*node).val, true
		}
	}

	return nil, false
}

func (sch *SCHash) Del(hk Key) {
	i := (hk.hashCode() % M)

	l := sch.Data[i]
	for e := sch.Data[i].Front(); e != nil; e = e.Next() {
		if e.Value.(*node).key == hk {
			l.Remove(e)
			if e.Next() != nil {
				l.MoveToBack(e.Next())
			}
		}
	}
}
