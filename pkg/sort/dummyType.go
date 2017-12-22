package sort

import (
	"math/rand"
	"time"
)

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

func (dn DummyNums) Shuffle() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < dn.Len(); i++ {
		j := rand.Intn(i + 1)
		dn.Swap(i, j)
	}
}

func NewRandNums(n int) DummyNums {
	rand.Seed(time.Now().Unix())
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
