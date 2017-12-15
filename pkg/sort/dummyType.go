package sort

import "math/rand"

type DummyNums []int

func (dn DummyNums) Len() int {
	return len(dn)
}

func (dn DummyNums) Less(i, j int) bool {
	return dn[i] < dn[j]
}

func (dn DummyNums) Swap(i, j int) {
	dn[i], dn[j] = dn[j], dn[i]
}

func NewRandNums(n int) DummyNums {
	dn := make(DummyNums, n)
	for i := 0; i < n; i++ {
		dn[i] = i
	}

	for i := 0; i < n; i++ {
		r := rand.Intn(i + 1)
		dn.Swap(i, r)
	}
	return dn
}

func CopyRandNums(rn DummyNums) DummyNums {
	dn := make(DummyNums, len(rn))
	for i, v := range rn {
		dn[i] = v
	}
	return dn
}
