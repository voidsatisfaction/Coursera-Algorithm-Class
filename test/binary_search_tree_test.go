package test

import (
	"reflect"
	"testing"

	"../pkg/binary_search_tree"
	"./dummy"
)

// type TreeKey int
//
// func (n TreeKey) Less(a interface{}) bool {
// 	return int(n) < int(a.(TreeKey))
// }
//
// func TestGet(t *testing.T) {
// 	bst := binarySearchTree.New()
// 	bst.Insert(TreeKey(8), 1)
// 	bst.Insert(TreeKey(3), 2)
// 	bst.Insert(TreeKey(5), 3)
// 	bst.Insert(TreeKey(6), 4)
//
// 	fmt.Printf("%+v", bst.Get(TreeKey(5)))
// }

func TestGet(t *testing.T) {
	bst := binarySearchTree.New()
	bst.Insert(dummy.NewTreeKey(7), 1)
	bst.Insert(dummy.NewTreeKey(9), 9)
	bst.Insert(dummy.NewTreeKey(5), []int{1, 2, 3})
	bst.Insert(dummy.NewTreeKey(1), 1)
	bst.Insert(dummy.NewTreeKey(-1), -1)

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
		if !reflect.DeepEqual(bst.Get(key), expect) {
			t.Errorf("expect: %+v, actual: %+v\n", expect, bst.Get(key))
		}
	}
}

func TestInclude(t *testing.T) {
	bst := binarySearchTree.New()
	bst.Insert(dummy.NewTreeKey(7), 1)
	bst.Insert(dummy.NewTreeKey(9), 9)
	bst.Insert(dummy.NewTreeKey(5), []int{1, 2, 3})
	bst.Insert(dummy.NewTreeKey(1), 1)
	bst.Insert(dummy.NewTreeKey(-1), -1)

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
		if !reflect.DeepEqual(bst.Include(key), expect) {
			t.Errorf("key: %+v, expect: %+v, actual: %+v\n", key, expect, bst.Include(key))
		}
	}
}
