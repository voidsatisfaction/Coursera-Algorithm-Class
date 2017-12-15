package sort

import (
	"reflect"
	"sort"
	"testing"
)

const (
	SIZE = 10000
)

func TestSelectionSort(t *testing.T) {
	dn1 := NewRandNums(SIZE)
	dn2 := CopyRandNums(dn1)
	sort.Ints(dn1)
	SelectionSort(dn2)
	if !reflect.DeepEqual(dn1, dn2) {
		t.Errorf("Expect: %+v, Got: %+v", dn1, dn2)
	}
}

func TestInsertionSort(t *testing.T) {
	dn1 := NewRandNums(SIZE)
	dn2 := CopyRandNums(dn1)
	sort.Ints(dn1)
	InsertionSort(dn2)
	if !reflect.DeepEqual(dn1, dn2) {
		t.Errorf("Expect: %+v, Got: %+v", dn1, dn2)
	}
}

func TestShellSort(t *testing.T) {
	dn1 := NewRandNums(SIZE)
	dn2 := CopyRandNums(dn1)
	sort.Ints(dn1)
	ShellSort(dn2)
	if !reflect.DeepEqual(dn1, dn2) {
		t.Errorf("Expect: %+v, Got: %+v", dn1, dn2)
	}
}
