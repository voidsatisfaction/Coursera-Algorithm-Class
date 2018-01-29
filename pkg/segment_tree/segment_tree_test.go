package segmentTree

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{0, 4, 2, 4, 1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{3, 7, 2, 5, 4, 1, 1, 9}, []int{0, 9, 7, 9, 7, 5, 4, 9, 3, 7, 2, 5, 4, 1, 1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}

	for _, test := range tests {
		arr, expected := test.arr, test.expected
		got := New(arr)
		if !reflect.DeepEqual(got, segmentTree(expected)) {
			t.Errorf("Segment tree is not equal")
			t.Errorf("Expect: %+v, got: %+v", expected, got)
		}
	}
}

func TestGet(t *testing.T) {
	arr := []int{3, 7, 2, 5, 4, 1, 1, 9}
	st, ed := 0, len(arr)-1
	tests := []struct {
		from     int
		to       int
		expected int
	}{
		{0, 7, 9},
		{-10, 10, 9},
		{-10, 2, 7},
		{-10, 5, 7},
		{1, 2, 7},
		{3, 10, 9},
	}

	segTree := New(arr)
	for _, test := range tests {
		from, to, expected := test.from, test.to, test.expected
		got := segTree.Get(st, ed, from, to, 1)
		if got != expected {
			t.Errorf("Range sum is not correct")
			t.Errorf("Expect: %d, got: %d", expected, got)
		}
	}
}

func TestUpdate(t *testing.T) {
	arr := []int{3, 7, 2, 5, 4, 1, 1, 9}
	st, ed := 0, len(arr)-1
	tests := []struct {
		i        int
		newVal   int
		expected []int
	}{
		{7, 3, []int{0, 7, 7, 4, 7, 5, 4, 3, 3, 7, 2, 5, 4, 1, 1, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{5, 5, []int{0, 9, 7, 9, 7, 5, 5, 9, 3, 7, 2, 5, 4, 5, 1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{1, 7, []int{0, 9, 7, 9, 7, 5, 4, 9, 3, 7, 2, 5, 4, 1, 1, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}

	for _, test := range tests {
		i, newVal, expected := test.i, test.newVal, test.expected
		segTree := New(arr)
		segTree.Update(st, ed, i, newVal, 1)
		if !reflect.DeepEqual(segTree, segmentTree(expected)) {
			t.Errorf("Updated segmentTree is not correct")
			t.Errorf("Expect: %+v, got: %+v", expected, segTree)
		}
	}
}
