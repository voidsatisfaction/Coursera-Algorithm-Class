package test

import (
	"reflect"
	"testing"

	"../pkg/redBlackTree"
	"./dummy"
)

func TestRedBlackTreeGet(t *testing.T) {
	rbt := redBlackTree.New()
	rbt.Insert(dummy.NewTreeKey(7), 1)
	rbt.Insert(dummy.NewTreeKey(9), 9)
	rbt.Insert(dummy.NewTreeKey(5), []int{1, 2, 3})
	rbt.Insert(dummy.NewTreeKey(1), 1)
	rbt.Insert(dummy.NewTreeKey(-1), -1)

	tests := []struct {
		key    dummy.TreeKey
		expect interface{}
	}{
		{dummy.NewTreeKey(5), []int{1, 2, 3}},
		{dummy.NewTreeKey(-100), nil},
		{dummy.NewTreeKey(9), 9},
	}

	for _, test := range tests {
		key, expect := test.key, test.expect
		if !reflect.DeepEqual(rbt.Get(key), expect) {
			t.Errorf("expect: %+v, actual: %+v\n", expect, rbt.Get(key))
		}
	}
}

func TestRedBlackTreeInclude(t *testing.T) {
	rbt := redBlackTree.New()
	rbt.Insert(dummy.NewTreeKey(7), 1)
	rbt.Insert(dummy.NewTreeKey(9), 9)
	rbt.Insert(dummy.NewTreeKey(5), []int{1, 2, 3})
	rbt.Insert(dummy.NewTreeKey(1), 1)
	rbt.Insert(dummy.NewTreeKey(-1), -1)

	tests := []struct {
		key    dummy.TreeKey
		expect bool
	}{
		{dummy.NewTreeKey(5), true},
		{dummy.NewTreeKey(-100), false},
		{dummy.NewTreeKey(9), true},
		{dummy.NewTreeKey(-1), true},
		{dummy.NewTreeKey(24), false},
	}

	for _, test := range tests {
		key, expect := test.key, test.expect
		if !reflect.DeepEqual(rbt.Include(key), expect) {
			t.Errorf("key: %+v, expect: %+v, actual: %+v\n", key, expect, rbt.Include(key))
		}
	}
}
