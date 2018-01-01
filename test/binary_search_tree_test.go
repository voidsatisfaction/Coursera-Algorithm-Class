package test

import (
	"reflect"
	"testing"

	"../pkg/binary_search_tree"
)

type TreeKey struct {
	key int
}

func (n TreeKey) Less(a interface{}) bool {
	return n.key < a.(TreeKey).key
}

func TestGet(t *testing.T) {
	bst := binarySearchTree.New()
	bst.Insert(TreeKey{7}, 1)
	bst.Insert(TreeKey{9}, 9)
	bst.Insert(TreeKey{5}, []int{1, 2, 3})
	bst.Insert(TreeKey{1}, 1)
	bst.Insert(TreeKey{-1}, -1)

	tests := []struct {
		key    TreeKey
		expect interface{}
	}{
		{TreeKey{5}, []int{1, 2, 3}},
		{TreeKey{-100}, nil},
		{TreeKey{9}, 9},
	}

	for _, test := range tests {
		key, expect := test.key, test.expect
		if !reflect.DeepEqual(bst.Get(key), expect) {
			t.Errorf("expect: %+v, actual: %+v\n", expect, bst.Get(key))
		}
	}
}

func TestInclude(t *testing.T) {
	bst := binarySearchTree.New()
	bst.Insert(TreeKey{7}, 1)
	bst.Insert(TreeKey{9}, 9)
	bst.Insert(TreeKey{5}, []int{1, 2, 3})
	bst.Insert(TreeKey{1}, 1)
	bst.Insert(TreeKey{-1}, -1)

	tests := []struct {
		key    TreeKey
		expect bool
	}{
		{TreeKey{5}, true},
		{TreeKey{-100}, false},
		{TreeKey{9}, true},
		{TreeKey{-1}, true},
		{TreeKey{24}, false},
	}

	for _, test := range tests {
		key, expect := test.key, test.expect
		if !reflect.DeepEqual(bst.Include(key), expect) {
			t.Errorf("key: %+v, expect: %+v, actual: %+v\n", key, expect, bst.Include(key))
		}
	}
}
